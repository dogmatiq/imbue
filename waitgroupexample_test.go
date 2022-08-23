package imbue_test

import (
	"context"
	"fmt"
	"time"

	"github.com/dogmatiq/imbue"
)

func ExampleContainer_WaitGroup() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type Dependency1 struct{ Value string }
	type Dependency2 struct{ Value string }

	// Declare a constructor for the Dependency1 type.
	imbue.With0(
		con,
		func(ctx imbue.Context) (Dependency1, error) {
			return Dependency1{"<value-1>"}, nil
		},
	)

	// Declare a constructor for the Dependency2 type.
	imbue.With0(
		con,
		func(ctx imbue.Context) (Dependency2, error) {
			return Dependency2{"<value-2>"}, nil
		},
	)

	// Create a wait group that is bound to the container.
	g := con.WaitGroup(context.Background())

	// Start some goroutines that depend on the dependencies.
	imbue.Go1(
		g,
		func(
			ctx context.Context,
			dep Dependency1,
		) error {
			fmt.Println(dep)
			return nil
		},
	)
	imbue.Go1(
		g,
		func(
			ctx context.Context,
			dep Dependency2,
		) error {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(dep)
			return nil
		},
	)

	// Wait for both goroutines to finish.
	if err := g.Wait(); err != nil {
		panic(err)
	}

	// Output:
	// {<value-1>}
	// {<value-2>}
}
