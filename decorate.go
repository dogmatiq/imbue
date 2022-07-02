package imbue

import "fmt"

// DecorateOption is an option that changes the behavior of a call to
// DecorateX().
type DecorateOption interface {
	applyDecorateOption()
}

// decorator is a container for a function that decorates a value of type T.
//
// It implements the userFunction interface.
type decorator[T any] struct {
	// impl is the decorator implementation. It is typically a closure generated
	// by the DecorateX() functions. It wraps the user-provided constructor
	// function to provide a common signature.
	impl func(*Context, T) (T, error)

	// loc is the location of the code that provided the decorator.
	loc location
}

// Call returns the decorated version of v.
func (d decorator[T]) Call(ctx *Context, v T) (T, error) {
	ctx = &Context{
		Context:  ctx,
		deferrer: ctx.deferrer,
		scope:    d,
	}

	v, err := d.impl(ctx, v)
	if err != nil {
		return v, fmt.Errorf(
			"%s failed: %w",
			d,
			err,
		)
	}

	return v, nil
}

// Location returns the location of the code that provided the decorator.
//
// This is typically the location of the call to the DecorateX() function, not
// the decorator implementation function definition.
func (d decorator[T]) Location() location {
	return d.loc
}

// String returns a description of the decorator for use in error messages.
func (d decorator[T]) String() string {
	return fmt.Sprintf(
		"%s decorator (%s)",
		typeOf[T](),
		d.loc,
	)
}

// Decorate adds a decorator function that is called after T's constructor.
func (d *declarationOf[T]) Decorate(
	impl func(*Context, T) (T, error),
	deps ...declaration,
) {
	dec := decorator[T]{
		impl,
		findLocation(),
	}

	for _, dep := range deps {
		d.dependsOn(dep, dec)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		panic(fmt.Sprintf(
			"cannot add %s because the value has already been constructed",
			dec,
		))
	}

	d.decorators = append(d.decorators, dec)
}

// decorate applies the declaration's decorators to d.value.
func (d *declarationOf[T]) decorate(ctx *Context) error {
	for _, dec := range d.decorators {
		var err error
		d.value, err = dec.Call(ctx, d.value)
		if err != nil {
			return err
		}
	}

	return nil
}
