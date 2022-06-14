package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

func ExampleInvoke1() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type Dependency struct{ Value string }

	// Declare a constructor for the Dependency type.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (Dependency, error) {
			return Dependency{"<value>"}, nil
		},
	)

	// Invoke a function that depends on the two dependency types declared
	// above.
	if err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep Dependency,
		) error {
			fmt.Println(dep)
			return nil
		},
	); err != nil {
		panic(err)
	}

	// Output:
	// {<value>}
}

func ExampleInvoke2() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type Dependency1 struct{ Value string }
	type Dependency2 struct{ Value string }

	// Declare a constructor for the Dependency1 type.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (Dependency1, error) {
			return Dependency1{"<value-1>"}, nil
		},
	)

	// Declare a constructor for the Dependency2 type.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (Dependency2, error) {
			return Dependency2{"<value-2>"}, nil
		},
	)

	// Invoke a function that depends on the two dependency types declared
	// above.
	if err := imbue.Invoke2(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep1 Dependency1,
			dep2 Dependency2,
		) error {
			fmt.Println(dep1)
			fmt.Println(dep2)
			return nil
		},
	); err != nil {
		panic(err)
	}

	// Output:
	// {<value-1>}
	// {<value-2>}
}
