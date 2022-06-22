// Code generated by Imbue's build process. DO NOT EDIT.

package imbue

// With0Grouped describes how to construct grouped values of type T.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With0Grouped[G Group, T any](
	con *Container,
	ctor func(*Context) (T, error),
	options ...WithGroupedOption,
) {
	With0(
		con,
		func(ctx *Context) (FromGroup[G, T], error) {
			v, err := ctor(ctx)
			return inGroup[G](v), err
		},
	)
}

// With1Grouped describes how to construct grouped values of type T from a
// single dependency.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With1Grouped[G Group, T, D any](
	con *Container,
	ctor func(*Context, D) (T, error),
	options ...WithGroupedOption,
) {
	With1(
		con,
		func(ctx *Context, v1 D) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1)
			return inGroup[G](v), err
		},
	)
}

// With2Grouped describes how to construct grouped values of type T from 2
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With2Grouped[G Group, T, D1, D2 any](
	con *Container,
	ctor func(*Context, D1, D2) (T, error),
	options ...WithGroupedOption,
) {
	With2(
		con,
		func(ctx *Context, v1 D1, v2 D2) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2)
			return inGroup[G](v), err
		},
	)
}

// With3Grouped describes how to construct grouped values of type T from 3
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With3Grouped[G Group, T, D1, D2, D3 any](
	con *Container,
	ctor func(*Context, D1, D2, D3) (T, error),
	options ...WithGroupedOption,
) {
	With3(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3)
			return inGroup[G](v), err
		},
	)
}

// With4Grouped describes how to construct grouped values of type T from 4
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With4Grouped[G Group, T, D1, D2, D3, D4 any](
	con *Container,
	ctor func(*Context, D1, D2, D3, D4) (T, error),
	options ...WithGroupedOption,
) {
	With4(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3, v4 D4) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4)
			return inGroup[G](v), err
		},
	)
}

// With5Grouped describes how to construct grouped values of type T from 5
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With5Grouped[G Group, T, D1, D2, D3, D4, D5 any](
	con *Container,
	ctor func(*Context, D1, D2, D3, D4, D5) (T, error),
	options ...WithGroupedOption,
) {
	With5(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5)
			return inGroup[G](v), err
		},
	)
}

// With6Grouped describes how to construct grouped values of type T from 6
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With6Grouped[G Group, T, D1, D2, D3, D4, D5, D6 any](
	con *Container,
	ctor func(*Context, D1, D2, D3, D4, D5, D6) (T, error),
	options ...WithGroupedOption,
) {
	With6(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6)
			return inGroup[G](v), err
		},
	)
}

// With7Grouped describes how to construct grouped values of type T from 7
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With7Grouped[G Group, T, D1, D2, D3, D4, D5, D6, D7 any](
	con *Container,
	ctor func(*Context, D1, D2, D3, D4, D5, D6, D7) (T, error),
	options ...WithGroupedOption,
) {
	With7(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6, v7 D7) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6, v7)
			return inGroup[G](v), err
		},
	)
}

// With8Grouped describes how to construct grouped values of type T from 8
// dependencies.
//
// G is the group that contains the dependency.
// T is the type of the dependency.
func With8Grouped[G Group, T, D1, D2, D3, D4, D5, D6, D7, D8 any](
	con *Container,
	ctor func(*Context, D1, D2, D3, D4, D5, D6, D7, D8) (T, error),
	options ...WithGroupedOption,
) {
	With8(
		con,
		func(ctx *Context, v1 D1, v2 D2, v3 D3, v4 D4, v5 D5, v6 D6, v7 D7, v8 D8) (FromGroup[G, T], error) {
			v, err := ctor(ctx, v1, v2, v3, v4, v5, v6, v7, v8)
			return inGroup[G](v), err
		},
	)
}
