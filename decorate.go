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
	loc location
	xfn decorator[T]
}

func (e decoratorEntry[T]) Call(ctx *Context, v T) (T, error) {
	ctx = &Context{
		Context:  ctx,
		deferrer: ctx.deferrer,
		scope:    e,
	}

	v, err := e.xfn(ctx, v)
	if err != nil {
		return v, fmt.Errorf(
			"%s failed: %w",
			e,
			err,
		)
	}

	return v, nil
}

func (e decoratorEntry[T]) String() string {
	return fmt.Sprintf(
		"%s decorator (%s)",
		typeOf[T](),
		e.loc,
	)
}

// Decorate adds a decorator function that is called after T's constructor.
func (d *declarationOf[T]) Decorate(
	fn decorator[T],
	deps ...declaration,
) {
	e := decoratorEntry[T]{
		findLocation(),
		fn,
	}

	d.m.Lock()
	if !d.isDeclared {
		d.location = e.loc
	}
	d.m.Unlock()

	for _, dep := range deps {
		d.dependsOn(dep, e)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		panic(fmt.Sprintf(
			"cannot add %s because the value has already been constructed",
			e,
		))
	}

	d.decorators = append(d.decorators, e)
}

// decorate applies the declaration's decorators to d.value.
func (d *declarationOf[T]) decorate(ctx *Context) error {
	for _, e := range d.decorators {
		var err error
		d.value, err = e.Call(ctx, d.value)
		if err != nil {
			return err
		}
	}

	return nil
}
