// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// Decorate0 describes how to decorate values of type T after construction.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate0[T any](
	con *Container,
	dec func(*Context, T) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			return dec, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate1 describes how to decorate values of type T after construction using
// a single additional dependency.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate1[T, D any](
	con *Container,
	dec func(*Context, T, D) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate2 describes how to decorate values of type T after construction using
// 2 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate2[T, D1, D2 any](
	con *Container,
	dec func(*Context, T, D1, D2) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate3 describes how to decorate values of type T after construction using
// 3 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate3[T, D1, D2, D3 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate4 describes how to decorate values of type T after construction using
// 4 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate4[T, D1, D2, D3, D4 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3, D4) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddDecoratorDependency(d4); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v4, err := d4.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3, v4)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate5 describes how to decorate values of type T after construction using
// 5 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate5[T, D1, D2, D3, D4, D5 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3, D4, D5) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddDecoratorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddDecoratorDependency(d5); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v4, err := d4.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v5, err := d5.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3, v4, v5)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate6 describes how to decorate values of type T after construction using
// 6 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate6[T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3, D4, D5, D6) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddDecoratorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddDecoratorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddDecoratorDependency(d6); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v4, err := d4.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v5, err := d5.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v6, err := d6.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3, v4, v5, v6)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate7 describes how to decorate values of type T after construction using
// 7 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate7[T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddDecoratorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddDecoratorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddDecoratorDependency(d6); err != nil {
				return nil, err
			}

			d7 := get[D7](con)
			if err := t.AddDecoratorDependency(d7); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v4, err := d4.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v5, err := d5.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v6, err := d6.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v7, err := d7.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3, v4, v5, v6, v7)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}

// Decorate8 describes how to decorate values of type T after construction using
// 8 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate8[T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	dec func(*Context, T, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...DecorateOption,
) {
	t := get[T](con)

	if err := t.AddDecorator(
		func() (decorator[T], error) {
			d1 := get[D1](con)
			if err := t.AddDecoratorDependency(d1); err != nil {
				return nil, err
			}

			d2 := get[D2](con)
			if err := t.AddDecoratorDependency(d2); err != nil {
				return nil, err
			}

			d3 := get[D3](con)
			if err := t.AddDecoratorDependency(d3); err != nil {
				return nil, err
			}

			d4 := get[D4](con)
			if err := t.AddDecoratorDependency(d4); err != nil {
				return nil, err
			}

			d5 := get[D5](con)
			if err := t.AddDecoratorDependency(d5); err != nil {
				return nil, err
			}

			d6 := get[D6](con)
			if err := t.AddDecoratorDependency(d6); err != nil {
				return nil, err
			}

			d7 := get[D7](con)
			if err := t.AddDecoratorDependency(d7); err != nil {
				return nil, err
			}

			d8 := get[D8](con)
			if err := t.AddDecoratorDependency(d8); err != nil {
				return nil, err
			}

			return func(ctx *Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v3, err := d3.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v4, err := d4.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v5, err := d5.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v6, err := d6.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v7, err := d7.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v8, err := d8.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2, v3, v4, v5, v6, v7, v8)
			}, nil
		},
	); err != nil {
		panic(err)
	}
}
