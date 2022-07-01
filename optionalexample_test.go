package imbue_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/dogmatiq/imbue"
)

func ExampleOptional() {
	con := imbue.New()
	defer con.Close()

	// Declare a type to use as a dependency within the example.
	type Dependency struct {
		Value string
	}

	// Declare a constructor for Dependency, but have it return an error.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (Dependency, error) {
			return Dependency{"<value>"}, nil
		},
	)

	// Invoke a function that optionally depends on the Dependency type.
	if err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep imbue.Optional[Dependency],
		) error {
			v, err := dep.Value()
			if err != nil {
				fmt.Println("dependency is unavailable:", err)
			} else {
				fmt.Println("dependency is available:", v)
			}

			return nil
		},
	); err != nil {
		panic(err)
	}
	// Output:
	// dependency is available: {<value>}
}

func ExampleOptional_failingConstructor() {
	con := imbue.New()
	defer con.Close()

	// Declare a type to use as a dependency within the example.
	type Dependency struct {
		Value string
	}

	// Declare a constructor for Dependency, but have it return an error.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (Dependency, error) {
			return Dependency{}, errors.New("<error>")
		},
	)

	// Invoke a function that optionally depends on the Dependency type.
	if err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep imbue.Optional[Dependency],
		) error {
			v, err := dep.Value()
			if err != nil {
				fmt.Println("dependency is unavailable:", err)
			} else {
				fmt.Println("dependency is available:", v)
			}

			return nil
		},
	); err != nil {
		panic(err)
	}
	// Output:
	// dependency is unavailable: imbue_test.Dependency constructor (optionalexample_test.go:62) failed: <error>
}

func ExampleOptional_constructorNotDeclared() {
	con := imbue.New()
	defer con.Close()

	// Declare a type to use as a dependency within the example.
	// Note that we don't actually declare a constructor for this type.
	type Dependency struct {
		Value string
	}

	// Invoke a function that optionally depends on the Dependency type.
	if err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			dep imbue.Optional[Dependency],
		) error {
			v, err := dep.Value()
			if err != nil {
				fmt.Println("dependency is unavailable:", err)
			} else {
				fmt.Println("dependency is available:", v)
			}

			return nil
		},
	); err != nil {
		panic(err)
	}
	// Output:
	// dependency is unavailable: no constructor is declared for imbue_test.Dependency
}
