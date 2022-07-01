package imbue

import (
	"errors"
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

	// dependsOn returns true if d depends on t, whether directly or indirectly.
	dependsOn(t declaration, cycle []declaration) ([]declaration, bool)

	// markAsDependency marks the declaration as a dependency. That is, other
	// declarations depend upon this one.
	markAsDependency()
}

// declarationOf describes how to build values of type T.
type declarationOf[T any] struct {
	m               sync.Mutex
	location        location
	isSelfDeclaring bool
	isConstructed   bool
	deps            map[reflect.Type]declaration
	isDep           bool
	constructor     constructor[T]
	decorators      []decoratorEntry[T]
	value           T
}

// selfDeclaring is an interface for types that construct themselves without a
// user-defined constructor function.
type selfDeclaring[T any] interface {
	declare(con *Container, decl *declarationOf[T]) error
}

// Init initializes the declaration.
func (d *declarationOf[T]) Init(con *Container) error {
	if sc, ok := any(d.value).(selfDeclaring[T]); ok {
		d.m.Lock()
		d.isSelfDeclaring = true
		d.m.Unlock()

		if err := sc.declare(con, d); err != nil {
			return err
		}
	}

	return nil
}

func (d *declarationOf[T]) addDependency(t declaration, funcType string) error {
	if cycle, ok := t.dependsOn(d, nil); ok {
		if len(cycle) == 1 {
			loc := findLocation()

			return fmt.Errorf(
				"%s for %s (%s) depends on itself",
				funcType,
				d.Type(),
				loc,
			)
		}

		message := fmt.Sprintf(
			"%s for %s introduces a cyclic dependency:",
			funcType,
			d.Type(),
		)

		for i := len(cycle) - 1; i >= 0; i-- {
			dep := cycle[i]
			loc := dep.BestLocation()

			message += fmt.Sprintf(
				"\n\t-> %s (%s)",
				dep.Type(),
				loc,
			)
		}

		return errors.New(message)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.deps == nil {
		d.deps = map[reflect.Type]declaration{}
	}

	d.deps[t.Type()] = t
	t.markAsDependency()

	return nil
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
	d.constructor = nil
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

	return d.location
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

// dependsOn returns true if d depends on t, whether directly or indirectly.
func (d *declarationOf[T]) dependsOn(t declaration, cycle []declaration) ([]declaration, bool) {
	if t.Type() == d.Type() {
		return append(cycle, d), true
	}

	d.m.Lock()
	defer d.m.Unlock()

	for _, dep := range d.deps {
		if cycle, ok := dep.dependsOn(t, cycle); ok {
			if d.isSelfDeclaring {
				return cycle, true
			}

			return append(cycle, d), true
		}
	}

	return nil, false
}

// markAsDependency marks the declaration as a dependency. That is, other
// declarations depend upon this one.
func (d *declarationOf[T]) markAsDependency() {
	d.m.Lock()
	defer d.m.Unlock()

	d.isDep = true
}
