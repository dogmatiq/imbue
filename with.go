package imbue

import (
	"fmt"
)

// WithOption is an option that changes the behavior of a call to WithX().
type WithOption interface {
	applyWithOption()
}

// constructor is a function that constructs values of type T.
type constructor[T any] func(*Context) (T, error)

type constructorEntry[T any] struct {
	loc    location
	impl   constructor[T]
	rawErr bool
}

func (e constructorEntry[T]) Call(ctx *Context) (T, error) {
	ctx = &Context{
		Context:  ctx,
		deferrer: ctx.deferrer,
		scope:    e,
	}

	v, err := e.impl(ctx)
	if err != nil {
		if e.rawErr {
			return v, err
		}

		return v, fmt.Errorf(
			"%s failed: %w",
			e,
			err,
		)
	}

	return v, nil
}

func (e constructorEntry[T]) String() string {
	return fmt.Sprintf(
		"%s constructor (%s)",
		typeOf[T](),
		e.loc,
	)
}

// Declare declares a constructor for values of type T.
func (d *declarationOf[T]) Declare(
	fn constructor[T],
	deps ...declaration,
) {
	e := constructorEntry[T]{
		findLocation(),
		fn,
		d.isSelfDeclaring,
	}

	d.m.Lock()
	d.location = e.loc
	d.m.Unlock()

	for _, dep := range deps {
		d.dependsOn(dep, e)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.isDeclared {
		isSelfDeclaring := d.isSelfDeclaring

		if isSelfDeclaring {
			panic(fmt.Sprintf(
				"explicit declaration of %s is disallowed",
				e,
			))
		}

		panic(fmt.Sprintf(
			"%s collides with existing constructor declared at %s",
			e,
			d.location,
		))
	}

	d.isDeclared = true
	d.constructor = e
}

// construct initializes d.value.
func (d *declarationOf[T]) construct(ctx *Context) error {
	if !d.isDeclared {
		return undeclaredConstructorError{d}
	}

	v, err := d.constructor.Call(ctx)
	if err != nil {
		return err
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
