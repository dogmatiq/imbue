package imbue_test

import (
	"fmt"

	"github.com/dogmatiq/imbue"
)

func ExampleDecorate1() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency struct{}
	type Dependency struct {
		Up *UpstreamDependency
	}

	// Declare a decorator for the Dependency type. It depends on the upstream
	// dependency type (which is assumed to be declared elsewhere).
	imbue.Decorate1(
		con,
		func(
			ctx *imbue.Context,
			d *Dependency,
			up *UpstreamDependency,
		) error {
			d.Up = up
			return nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *imbue_test.Dependency
	//     └── *imbue_test.UpstreamDependency
}

func ExampleDecorate2() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency1 struct{}
	type UpstreamDependency2 struct{}
	type Dependency struct {
		Up1 *UpstreamDependency1
		Up2 *UpstreamDependency2
	}

	// Declare a decorator for the Dependency type. It depends on the two
	// upstream dependency types (which are assumed to be declared elsewhere).
	imbue.Decorate2(
		con,
		func(
			ctx *imbue.Context,
			d *Dependency,
			up1 *UpstreamDependency1,
			up2 *UpstreamDependency2,
		) error {
			d.Up1 = up1
			d.Up2 = up2
			return nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *imbue_test.Dependency
	//     ├── *imbue_test.UpstreamDependency1
	//     └── *imbue_test.UpstreamDependency2
}
