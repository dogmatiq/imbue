package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

func ExampleCatalog() {
	cat := imbue.NewCatalog()

	// Declare a type to use as a dependency within the example.
	type Dependency struct {
		Count int
	}

	// Keep track of how many times we have constructed the dependency.
	count := 0

	// Declare a constructor for the Dependency type within the catalog.
	imbue.With0(
		cat,
		func(
			ctx imbue.Context,
		) (*Dependency, error) {
			count++
			return &Dependency{
				Count: count,
			}, nil
		},
	)

	// Create a new container from the catalog.
	con1 := imbue.New(imbue.WithCatalog(cat))
	defer con1.Close()

	imbue.Invoke1(
		context.Background(),
		con1,
		func(
			ctx context.Context,
			d *Dependency,
		) error {
			fmt.Println("count is", d.Count)
			return nil
		},
	)

	// Create a second container from the same catalog.
	con2 := imbue.New(imbue.WithCatalog(cat))
	defer con2.Close()

	imbue.Invoke1(
		context.Background(),
		con2,
		func(
			ctx context.Context,
			d *Dependency,
		) error {
			fmt.Println("count is", d.Count)
			return nil
		},
	)

	// Output:
	// count is 1
	// count is 2
}
