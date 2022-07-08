package imbue

import (
	"context"
	"fmt"
)

// DecorateOption is an option that changes the behavior of a call to
// DecorateX().
type DecorateOption interface {
	applyDecorateOption()
}

// decorator is a wrapper around a function that decorates a value of type T.
//
// It implements the userFunction interface.
type decorator[T any] struct {
	// impl is the decorator implementation. It is typically a closure generated
	// by the DecorateX() functions. It wraps the user-provided decorator
	// function to provide a common signature.
	impl func(Context, T) (T, error)

	// loc is the location of the code that provided the decorator.
	loc location
}

// Call returns the decorated version of v.
func (d decorator[T]) Call(ctx context.Context, v T, defers *deferSet) (T, error) {
	v, err := d.impl(
		&scopedContext{
			Context: ctx,
			scope:   d,
			defers:  defers,
		},
		v,
	)
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
