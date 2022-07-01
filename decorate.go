package imbue

import "fmt"

// DecorateOption is an option that changes the behavior of a call to
// DecorateX().
type DecorateOption interface {
	applyDecorateOption()
}

// decorator is a function that is called after T's constructor.
type decorator[T any] func(*Context, T) (T, error)

// decoratorEntry encapsulates a decorator and information about where it was
// declared.
type decoratorEntry[T any] struct {
	Location  location
	Decorator decorator[T]
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
			d.Type(),
			loc,
		)
	}

	d.decorators = append(d.decorators, e)

	return nil
}

// decorate applies the decorators to d.value.
func (d *declarationOf[T]) decorate(ctx *Context) error {
	for _, e := range d.decorators {
		var err error
		d.value, err = e.Decorator(
			ctx.newChild("decorator", d.Type()),
			d.value,
		)
		if err != nil {
			return fmt.Errorf(
				"decorator for %s (%s) failed: %w",
				d.Type(),
				e.Location,
				err,
			)
		}
	}

	return nil
}
