package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

// Closable is a dependency that can be closed.
type Closable struct{}

func (c Closable) Close() error {
	fmt.Println("closed!")
	return nil
}

func ExampleContext_Defer() {
	con := imbue.New()
	defer con.Close()

	// Declare a constructor for a Closable dependency.
	imbue.With0(
		con,
		func(
			ctx *imbue.Context,
		) (Closable, error) {
			c := Closable{}
			ctx.Defer(c.Close)

			return c, nil
		},
	)

	// Invoke a function that depends on both the Closable dependency.
	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			c Closable,
		) error {
			fmt.Println("invoked!")
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
	// Output:
	// invoked!
	// closed!
}
