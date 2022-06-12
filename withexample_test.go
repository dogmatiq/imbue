package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

func ExampleWith0() {
	con := imbue.New()

	type SomeDependency struct {
		Value string
	}

	// Teach the container how to construct SomeDependency values.
	imbue.With0(
		con,
		func(
			ctx *imbue.Context,
		) (SomeDependency, error) {
			return SomeDependency{Value: "<value>"}, nil
		},
	)

	// Invoke a function that depends on the SomeDependency value constructed by
	// the container.
	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep SomeDependency,
		) error {
			fmt.Println(dep.Value)
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	// Output:
	// <value>
}
