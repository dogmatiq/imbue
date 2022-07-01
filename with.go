package imbue

import (
	"errors"
	"fmt"
)

// WithOption is an option that changes the behavior of a call to WithX().
type WithOption interface {
	applyWithOption()
}

// constructor is a function that constructs values of type T.
type constructor[T any] func(*Context) (T, error)

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
				d.Type(),
				loc,
			)
		}

		return fmt.Errorf(
			"constructor for %s (%s) collides with existing constructor declared at %s",
			d.Type(),
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

// AddConstructorDependency marks t as a dependency of d's constructor.
func (d *declarationOf[T]) AddConstructorDependency(t declaration) error {
	return d.addDependency(t, "constructor")
}

// construct initializes d.value.
func (d *declarationOf[T]) construct(ctx *Context) error {
	if d.constructor == nil {
		return undeclaredConstructorError{d}
	}

	v, err := d.constructor(
		ctx.newChild("constructor", d.Type()),
	)
	if err != nil {
		// If the type is self-declaring let it specify the exact error.
		if d.isSelfDeclaring {
			return err
		}

		// Otherwise, wrap the error with file/line information.
		return fmt.Errorf(
			"constructor for %s (%s) failed: %w",
			d.Type(),
			d.location,
			err,
		)
	}

	d.value = v

	return nil
}

// undeclaredConstructorError is an error returned by declarationOf[T].Resolve()
// when no constructor has been declared for T.
type undeclaredConstructorError struct {
	Declaration declaration
}

func (e undeclaredConstructorError) Error() string {
	return fmt.Sprintf(
		"no constructor is declared for %s",
		e.Declaration.Type(),
	)
}

// panicOnUndeclaredConstructor panics if err is an undeclaredConstructor error.
func panicOnUndeclaredConstructor(err error) {
	var u undeclaredConstructorError
	if errors.As(err, &u) {
		panic(u)
	}
}
