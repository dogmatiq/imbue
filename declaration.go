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
	// GetType returns the type of the value constructed by this declaration.
	GetType() reflect.Type

	// Location returns the location of the declaration in code.
	Location() location

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

// constructor is a function that constructs values of type T.
type constructor[T any] func(*Context) (T, error)

// decorator is a function that is called after T's constructor.
type decorator[T any] func(*Context, T) (T, error)

// decoratorEntry encapsulates a decorator and information about where it was
// declared.
type decoratorEntry[T any] struct {
	Location  location
	Decorator decorator[T]
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

// Declare declares a constructor for values of type T.
func (d *declarationOf[T]) Declare(
	decl func() (constructor[T], error),
) error {
	loc := findLocation()

	d.m.Lock()

	if d.constructor != nil {
		isSelfDeclaring := d.isSelfDeclaring
		d.m.Unlock()

		if isSelfDeclaring {
			return fmt.Errorf(
				"explicit declaration of constructor for %s (%s) is disallowed",
				d.GetType(),
				loc,
			)
		}

		return fmt.Errorf(
			"constructor for %s (%s) collides with existing constructor declared at %s",
			d.GetType(),
			loc,
			d.location,
		)
	}

	d.location = loc

	d.m.Unlock()

	c, err := decl()
	if err != nil {
		return err
	}

	d.m.Lock()
	defer d.m.Unlock()

	d.constructor = c

	return nil
}

// AddDecorator adds a decorator function that is called after T's
// constructor.
func (d *declarationOf[T]) AddDecorator(
	decl func() (decorator[T], error),
) error {
	loc := findLocation()

	d.m.Lock()
	if d.constructor == nil {
		d.location = loc
	}
	d.m.Unlock()

	i, err := decl()
	if err != nil {
		return err
	}

	e := decoratorEntry[T]{
		Location:  loc,
		Decorator: i,
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		return fmt.Errorf(
			"cannot add decorator for %s (%s) because it has already been constructed",
			d.GetType(),
			loc,
		)
	}

	d.decorators = append(d.decorators, e)

	return nil
}

// AddConstructorDependency marks t as a dependency of d's constructor.
func (d *declarationOf[T]) AddConstructorDependency(t declaration) error {
	return d.addDependency(t, "constructor")
}

// AddDecoratorDependency marks t as a dependency of one of d's decorators.
func (d *declarationOf[T]) AddDecoratorDependency(t declaration) error {
	return d.addDependency(t, "decorator")
}

func (d *declarationOf[T]) addDependency(t declaration, funcType string) error {
	if cycle, ok := t.dependsOn(d, nil); ok {
		if len(cycle) == 1 {
			loc := findLocation()

			return fmt.Errorf(
				"%s for %s (%s) depends on itself",
				funcType,
				d.GetType(),
				loc,
			)
		}

		message := fmt.Sprintf(
			"%s for %s introduces a cyclic dependency:",
			funcType,
			d.GetType(),
		)

		for i := len(cycle) - 1; i >= 0; i-- {
			dep := cycle[i]
			loc := dep.Location()

			message += fmt.Sprintf(
				"\n\t-> %s (%s)",
				dep.GetType(),
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

	d.deps[t.GetType()] = t
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

// construct initializes d.value.
func (d *declarationOf[T]) construct(ctx *Context) error {
	if d.constructor == nil {
		return undeclaredConstructor{d}
	}

	v, err := d.constructor(ctx)
	if err != nil {
		// If the type is self-declaring let it specify the exact error.
		if d.isSelfDeclaring {
			return err
		}

		// Otherwise, wrap the error with file/line information.
		return fmt.Errorf(
			"constructor for %s (%s) failed: %w",
			d.GetType(),
			d.location,
			err,
		)
	}

	d.value = v

	return nil
}

// decorate applies the decorators to d.value.
func (d *declarationOf[T]) decorate(ctx *Context) error {
	for _, e := range d.decorators {
		var err error
		d.value, err = e.Decorator(ctx, d.value)
		if err != nil {
			return fmt.Errorf(
				"decorator for %s (%s) failed: %w",
				d.GetType(),
				e.Location,
				err,
			)
		}
	}

	return nil
}

// GetType returns the type of the value constructed by this declaration.
func (d *declarationOf[T]) GetType() reflect.Type {
	return typeOf[T]()
}

// Location returns the location of the declaration in code.
func (d *declarationOf[T]) Location() location {
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
	if t.GetType() == d.GetType() {
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

// undeclaredConstructor is an error returned by declarationOf[T].Resolve() when
// no constructor has been declared for T.
type undeclaredConstructor struct {
	Declaration declaration
}

func (e undeclaredConstructor) Error() string {
	return fmt.Sprintf(
		"no constructor is declared for %s",
		e.Declaration.GetType(),
	)
}

// panicOnUndeclaredConstructor panics if err is an undeclaredConstructor error.
func panicOnUndeclaredConstructor(err error) {
	var u undeclaredConstructor
	if errors.As(err, &u) {
		panic(u)
	}
}
