// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// Decorate0 describes how to decorate values of type T after construction.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate0[T any](
	con ContainerAware,
	dec func(Context, T) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
				return dec(ctx, v)
			},
		)
	})
}

// Decorate1 describes how to decorate values of type T after construction using
// a single additional dependency.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate1[T, D any](
	con ContainerAware,
	dec func(Context, T, D) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1)
			},
			d1,
		)
	})
}

// Decorate2 describes how to decorate values of type T after construction using
// 2 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate2[T, D1, D2 any](
	con ContainerAware,
	dec func(Context, T, D1, D2) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
				v1, err := d1.Resolve(ctx)
				if err != nil {
					return v, err
				}

				v2, err := d2.Resolve(ctx)
				if err != nil {
					return v, err
				}

				return dec(ctx, v, v1, v2)
			},
			d1,
			d2,
		)
	})
}

// Decorate3 describes how to decorate values of type T after construction using
// 3 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate3[T, D1, D2, D3 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
		)
	})
}

// Decorate4 describes how to decorate values of type T after construction using
// 4 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate4[T, D1, D2, D3, D4 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3, D4) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)
		d4 := get[D4](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
			d4,
		)
	})
}

// Decorate5 describes how to decorate values of type T after construction using
// 5 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate5[T, D1, D2, D3, D4, D5 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3, D4, D5) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)
		d4 := get[D4](con)
		d5 := get[D5](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
			d4,
			d5,
		)
	})
}

// Decorate6 describes how to decorate values of type T after construction using
// 6 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate6[T, D1, D2, D3, D4, D5, D6 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3, D4, D5, D6) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)
		d4 := get[D4](con)
		d5 := get[D5](con)
		d6 := get[D6](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
			d4,
			d5,
			d6,
		)
	})
}

// Decorate7 describes how to decorate values of type T after construction using
// 7 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate7[T, D1, D2, D3, D4, D5, D6, D7 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)
		d4 := get[D4](con)
		d5 := get[D5](con)
		d6 := get[D6](con)
		d7 := get[D7](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
			d4,
			d5,
			d6,
			d7,
		)
	})
}

// Decorate8 describes how to decorate values of type T after construction using
// 8 additional dependencies.
//
// The dependency being decorated is passed to dec and replaced with
// the decorator's return value.
//
// The decorated dependency may be manipulated in-place.
func Decorate8[T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con ContainerAware,
	dec func(Context, T, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...DecorateOption,
) {
	con.withContainer(func(con *Container) {
		t := get[T](con)
		d1 := get[D1](con)
		d2 := get[D2](con)
		d3 := get[D3](con)
		d4 := get[D4](con)
		d5 := get[D5](con)
		d6 := get[D6](con)
		d7 := get[D7](con)
		d8 := get[D8](con)

		t.Decorate(
			func(ctx Context, v T) (T, error) {
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
			},
			d1,
			d2,
			d3,
			d4,
			d5,
			d6,
			d7,
			d8,
		)
	})
}
