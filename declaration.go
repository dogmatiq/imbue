package imbue

import (
	"fmt"
	"reflect"
	"sync"
)

// declaration is an interface that describes how to build a value of a specific
// type.
type declaration interface {
	// Type returns the type of the value constructed by this declaration.
	Type() reflect.Type

	// BestLocation returns the "best" known location of the declaration in
	// code.
	//
	// Typically this is the location of the constructor for the definition, but
	// it may refer to some other location (such as a decorator function) if the
	// constructor has not yet been defined.
	BestLocation() location

	// IsDependency returns true if other declarations depend upon this one.
	IsDependency() bool

	// Dependencies returns the declarations that this declaration depends upon,
	// sorted by type.
	Dependencies() []declaration

	// IsImplicit returns true if this is an implicit declaration.
	//
	// Implicit declarations are added to the container as needed without the
	// user declaring a constructor function.
	IsImplicit() bool

	// MarkAsDependency marks the declaration as a dependency. That is, other
	// declarations depend upon this one.
	MarkAsDependency()
}

// findPath returns the path from t to d, where d is a (possibly indirect)
// dependency of t.
//
// If t does not depend on d, the path is an empty slice.
//
// If t is d, the path is a single-element slice containing t.
func findPath(t, d declaration) []declaration {
	if t == d {
		return []declaration{t}
	}

	for _, dep := range t.Dependencies() {
		if p := findPath(dep, d); len(p) != 0 {
			if t.IsImplicit() {
				return p
			}

			return append(p, t)
		}
	}

	return nil
}

// declarationOf describes how to build values of type T.
type declarationOf[T any] struct {
	m               sync.Mutex
	initLocation    location
	isSelfDeclaring bool
	isDeclared      bool
	isConstructed   bool
	deps            map[reflect.Type]declaration
	isDep           bool
	constructor     constructor[T]
	decorators      []decorator[T]
	value           T
}

// userFunction is an interface for a user-supplied function that forms part of
// the life-cycle of a specific type, such as constructors and decorators.
type userFunction interface {
	Location() location
	String() string
}

// selfDeclaring is an interface for types that construct themselves without a
// user-defined constructor function.
type selfDeclaring[T any] interface {
	declare(con *Container, decl *declarationOf[T])
}

// Init initializes the declaration.
//
// It is called when the declaration is first added to the container.
func (d *declarationOf[T]) Init(con *Container) {
	sc, ok := any(d.value).(selfDeclaring[T])

	d.m.Lock()
	d.initLocation = findLocation()
	d.isSelfDeclaring = ok
	d.m.Unlock()

	if ok {
		sc.declare(con, d)
	}
}

// dependsOn adds a dependency on type t.
func (d *declarationOf[T]) dependsOn(t declaration, scope userFunction) {
	path := findPath(t, d)

	if len(path) == 1 {
		panic(fmt.Sprintf(
			"%s depends on itself",
			scope,
		))
	}

	if len(path) != 0 {
		message := fmt.Sprintf(
			"%s introduces a cyclic dependency:",
			scope,
		)

		for i := len(path) - 1; i >= 0; i-- {
			dep := path[i]

			message += fmt.Sprintf(
				"\n\t-> %s (%s)",
				dep.Type(),
				dep.BestLocation(),
			)
		}

		panic(message)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.deps == nil {
		d.deps = map[reflect.Type]declaration{}
	}

	d.deps[t.Type()] = t
	t.MarkAsDependency()
}

// Resolve returns the value constructed by this declaration.
//
// The constructor is called only once. Subsequent calls to Resolve() return the
// same value.
func (d *declarationOf[T]) Resolve(ctx *Context) (T, error) {
	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		return d.value, nil
	}

	if err := d.construct(ctx); err != nil {
		return d.value, err
	}

	if err := d.decorate(ctx); err != nil {
		return d.value, err
	}

	d.isConstructed = true
	d.constructor = constructor[T]{}
	d.decorators = nil

	return d.value, nil
}

// Type returns the type of the value constructed by this declaration.
func (d *declarationOf[T]) Type() reflect.Type {
	return typeOf[T]()
}

// BestLocation returns the "best" known location of the declaration in code.
//
// Typically this is the location of the constructor for the definition, but it
// may refer to some other location (such as a decorator function) if the
// constructor has not yet been defined.
func (d *declarationOf[T]) BestLocation() location {
	d.m.Lock()
	defer d.m.Unlock()

	if d.isDeclared {
		return d.constructor.Location()
	}

	for _, dec := range d.decorators {
		return dec.Location()
	}

	return d.initLocation
}

// IsDependency returns true if other declarations depend upon this one.
func (d *declarationOf[T]) IsDependency() bool {
	d.m.Lock()
	defer d.m.Unlock()

	return d.isDep
}

// Dependencies returns the declarations that this declaration depends upon,
// sorted by type.
func (d *declarationOf[T]) Dependencies() []declaration {
	d.m.Lock()
	defer d.m.Unlock()

	return sortDeclarations(d.deps)
}

// IsImplicit returns true if this is an implicit declaration.
//
// Implicit declarations are added to the container as needed without the
// user declaring a constructor function.
func (d *declarationOf[T]) IsImplicit() bool {
	d.m.Lock()
	defer d.m.Unlock()

	return d.isSelfDeclaring
}

// MarkAsDependency marks the declaration as a dependency. That is, other
// declarations depend upon this one.
func (d *declarationOf[T]) MarkAsDependency() {
	d.m.Lock()
	defer d.m.Unlock()

	d.isDep = true
}
