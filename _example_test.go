package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

type Dependency1 string

func (Dependency1) Close() error {
	fmt.Println("closed dependency")
	return nil
}

type Dependency2 string

type SomeType struct {
	Dep1         Dependency1
	Dep2A, Dep2B Dependency2
}

func Example() {
	con := imbue.New()

	imbue.With3(con, func(
		ctx *imbue.Context,
		meaningless Dependency1,
		red Dependency2,
		green Dependency2,
	) (SomeType, error) {
		return SomeType{
			Dep1:  meaningless,
			Dep2A: red,
			Dep2B: green,
		}, nil
	})

	imbue.With0(con, func(
		ctx *imbue.Context,
	) (Dependency1, error) {
		d := Dependency1("<dependency>")
		ctx.Defer(d.Close)

		return d, nil
	})

	imbue.With0(con, func(
		ctx *imbue.Context,
	) (red Dependency2, _ error) {
		return Dependency2("<red>"), nil
	})

	imbue.With0(con, func(
		ctx *imbue.Context,
	) (green Dependency2, _ error) {
		return Dependency2("<green>"), nil
	})

	v, err := imbue.CallWith1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			s SomeType,
		) (int, error) {
			fmt.Println(s.Dep1)
			fmt.Println(s.Dep2A)
			fmt.Println(s.Dep2B)
			return 23, nil
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Print("value: ", v)
	con.Close()

	// Output:
	// <dependency>
	// <red>
	// <green>
	// value: 23
	// closed dependency
}
