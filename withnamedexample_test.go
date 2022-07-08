package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

// Color is a type that represents a color.
type Color string

type (
	// Foreground is a name for a color.
	Foreground imbue.Name[Color]

	// Background is a name for a color.
	Background imbue.Name[Color]
)

func ExampleByName() {
	con := imbue.New()
	defer con.Close()

	// Declare a constructor for a Color named Foreground.
	imbue.With0Named[Foreground](
		con,
		func(
			ctx imbue.Context,
		) (Color, error) {
			return "<black>", nil
		},
	)

	// Declare a constructor for a Color named Background.
	imbue.With0Named[Background](
		con,
		func(
			ctx imbue.Context,
		) (Color, error) {
			return "<white>", nil
		},
	)

	// Invoke a function that depends on both the Foreground and Background
	// colors.
	err := imbue.Invoke2(
		context.Background(),
		con,
		func(
			ctx context.Context,
			fg imbue.ByName[Foreground, Color],
			bg imbue.ByName[Background, Color],
		) error {
			// Named dependencies have a Name() and Value() methods which return
			// the name and value of the dependency.
			fmt.Println(fg.Name(), "=", fg.Value())
			fmt.Println(bg.Name(), "=", bg.Value())
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
	// Output:
	// Foreground = <black>
	// Background = <white>
}
