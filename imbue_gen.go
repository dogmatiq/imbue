// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

import "context"

// With0 describes how to construct values of type T.
func With0[T any](
	con *Container,
	fn func(*Context) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			return fn(ctx)
		},
	)
}

// With1 describes how to construct values of type T from a single dependency.
func With1[T, D any](
	con *Container,
	fn func(*Context, D) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep, err := resolve[D](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep)
		},
	)
}

// With2 describes how to construct values of type T from 2 dependencies.
func With2[T, D1, D2 any](
	con *Container,
	fn func(*Context, D1, D2) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2)
		},
	)
}

// With3 describes how to construct values of type T from 3 dependencies.
func With3[T, D1, D2, D3 any](
	con *Container,
	fn func(*Context, D1, D2, D3) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3)
		},
	)
}

// With4 describes how to construct values of type T from 4 dependencies.
func With4[T, D1, D2, D3, D4 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4)
		},
	)
}

// With5 describes how to construct values of type T from 5 dependencies.
func With5[T, D1, D2, D3, D4, D5 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5)
		},
	)
}

// With6 describes how to construct values of type T from 6 dependencies.
func With6[T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6)
		},
	)
}

// With7 describes how to construct values of type T from 7 dependencies.
func With7[T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7)
		},
	)
}

// With8 describes how to construct values of type T from 8 dependencies.
func With8[T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8)
		},
	)
}

// With9 describes how to construct values of type T from 9 dependencies.
func With9[T, D1, D2, D3, D4, D5, D6, D7, D8, D9 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9)
		},
	)
}

// With10 describes how to construct values of type T from 10 dependencies.
func With10[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10)
		},
	)
}

// With11 describes how to construct values of type T from 11 dependencies.
func With11[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			dep11, err := resolve[D11](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11)
		},
	)
}

// With12 describes how to construct values of type T from 12 dependencies.
func With12[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			dep11, err := resolve[D11](ctx, con)
			if err != nil {
				return value, err
			}

			dep12, err := resolve[D12](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12)
		},
	)
}

// With13 describes how to construct values of type T from 13 dependencies.
func With13[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			dep11, err := resolve[D11](ctx, con)
			if err != nil {
				return value, err
			}

			dep12, err := resolve[D12](ctx, con)
			if err != nil {
				return value, err
			}

			dep13, err := resolve[D13](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13)
		},
	)
}

// With14 describes how to construct values of type T from 14 dependencies.
func With14[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			dep11, err := resolve[D11](ctx, con)
			if err != nil {
				return value, err
			}

			dep12, err := resolve[D12](ctx, con)
			if err != nil {
				return value, err
			}

			dep13, err := resolve[D13](ctx, con)
			if err != nil {
				return value, err
			}

			dep14, err := resolve[D14](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13, dep14)
		},
	)
}

// With15 describes how to construct values of type T from 15 dependencies.
func With15[T, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14, D15 any](
	con *Container,
	fn func(*Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14, D15) (T, error),
) {
	register(
		con,
		func(ctx *Context, con *Container) (value T, _ error) {
			dep1, err := resolve[D1](ctx, con)
			if err != nil {
				return value, err
			}

			dep2, err := resolve[D2](ctx, con)
			if err != nil {
				return value, err
			}

			dep3, err := resolve[D3](ctx, con)
			if err != nil {
				return value, err
			}

			dep4, err := resolve[D4](ctx, con)
			if err != nil {
				return value, err
			}

			dep5, err := resolve[D5](ctx, con)
			if err != nil {
				return value, err
			}

			dep6, err := resolve[D6](ctx, con)
			if err != nil {
				return value, err
			}

			dep7, err := resolve[D7](ctx, con)
			if err != nil {
				return value, err
			}

			dep8, err := resolve[D8](ctx, con)
			if err != nil {
				return value, err
			}

			dep9, err := resolve[D9](ctx, con)
			if err != nil {
				return value, err
			}

			dep10, err := resolve[D10](ctx, con)
			if err != nil {
				return value, err
			}

			dep11, err := resolve[D11](ctx, con)
			if err != nil {
				return value, err
			}

			dep12, err := resolve[D12](ctx, con)
			if err != nil {
				return value, err
			}

			dep13, err := resolve[D13](ctx, con)
			if err != nil {
				return value, err
			}

			dep14, err := resolve[D14](ctx, con)
			if err != nil {
				return value, err
			}

			dep15, err := resolve[D15](ctx, con)
			if err != nil {
				return value, err
			}

			return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13, dep14, dep15)
		},
	)
}

// Invoke1 calls a function with a single dependency.
func Invoke1[D any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D) error,
) error {
	rctx := rootContext(ctx)

	dep, err := resolve[D](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep)
}

// Invoke2 calls a function with 2 dependencies.
func Invoke2[D1, D2 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2)
}

// Invoke3 calls a function with 3 dependencies.
func Invoke3[D1, D2, D3 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3)
}

// Invoke4 calls a function with 4 dependencies.
func Invoke4[D1, D2, D3, D4 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4)
}

// Invoke5 calls a function with 5 dependencies.
func Invoke5[D1, D2, D3, D4, D5 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5)
}

// Invoke6 calls a function with 6 dependencies.
func Invoke6[D1, D2, D3, D4, D5, D6 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6)
}

// Invoke7 calls a function with 7 dependencies.
func Invoke7[D1, D2, D3, D4, D5, D6, D7 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7)
}

// Invoke8 calls a function with 8 dependencies.
func Invoke8[D1, D2, D3, D4, D5, D6, D7, D8 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8)
}

// Invoke9 calls a function with 9 dependencies.
func Invoke9[D1, D2, D3, D4, D5, D6, D7, D8, D9 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9)
}

// Invoke10 calls a function with 10 dependencies.
func Invoke10[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10)
}

// Invoke11 calls a function with 11 dependencies.
func Invoke11[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	dep11, err := resolve[D11](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11)
}

// Invoke12 calls a function with 12 dependencies.
func Invoke12[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	dep11, err := resolve[D11](rctx, con)
	if err != nil {
		return err
	}

	dep12, err := resolve[D12](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12)
}

// Invoke13 calls a function with 13 dependencies.
func Invoke13[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	dep11, err := resolve[D11](rctx, con)
	if err != nil {
		return err
	}

	dep12, err := resolve[D12](rctx, con)
	if err != nil {
		return err
	}

	dep13, err := resolve[D13](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13)
}

// Invoke14 calls a function with 14 dependencies.
func Invoke14[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	dep11, err := resolve[D11](rctx, con)
	if err != nil {
		return err
	}

	dep12, err := resolve[D12](rctx, con)
	if err != nil {
		return err
	}

	dep13, err := resolve[D13](rctx, con)
	if err != nil {
		return err
	}

	dep14, err := resolve[D14](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13, dep14)
}

// Invoke15 calls a function with 15 dependencies.
func Invoke15[D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14, D15 any](
	ctx context.Context,
	con *Container,
	fn func(context.Context, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10, D11, D12, D13, D14, D15) error,
) error {
	rctx := rootContext(ctx)

	dep1, err := resolve[D1](rctx, con)
	if err != nil {
		return err
	}

	dep2, err := resolve[D2](rctx, con)
	if err != nil {
		return err
	}

	dep3, err := resolve[D3](rctx, con)
	if err != nil {
		return err
	}

	dep4, err := resolve[D4](rctx, con)
	if err != nil {
		return err
	}

	dep5, err := resolve[D5](rctx, con)
	if err != nil {
		return err
	}

	dep6, err := resolve[D6](rctx, con)
	if err != nil {
		return err
	}

	dep7, err := resolve[D7](rctx, con)
	if err != nil {
		return err
	}

	dep8, err := resolve[D8](rctx, con)
	if err != nil {
		return err
	}

	dep9, err := resolve[D9](rctx, con)
	if err != nil {
		return err
	}

	dep10, err := resolve[D10](rctx, con)
	if err != nil {
		return err
	}

	dep11, err := resolve[D11](rctx, con)
	if err != nil {
		return err
	}

	dep12, err := resolve[D12](rctx, con)
	if err != nil {
		return err
	}

	dep13, err := resolve[D13](rctx, con)
	if err != nil {
		return err
	}

	dep14, err := resolve[D14](rctx, con)
	if err != nil {
		return err
	}

	dep15, err := resolve[D15](rctx, con)
	if err != nil {
		return err
	}

	return fn(ctx, dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10, dep11, dep12, dep13, dep14, dep15)
}
