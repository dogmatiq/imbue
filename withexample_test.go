package imbue_test

import (
	"github.com/dogmatiq/imbue"
)

func ExampleWith0() {
	con := imbue.New()
	defer con.Close()

	// Declare a type to use as a dependency within the example.
	type Dependency struct{}

	// Declare a constructor for the Dependency type.
	imbue.With0(
		con,
		func(ctx *imbue.Context) (*Dependency, error) {
			return &Dependency{}, nil
		},
	)
}

func ExampleWith1() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency struct{}
	type Dependency struct {
		Up *UpstreamDependency
	}

	// Declare a constructor for the Dependency type. It depends on the upstream
	// dependency type (which is assumed to be declared elsewhere).
	imbue.With1(
		con,
		func(
			ctx *imbue.Context,
			up *UpstreamDependency,
		) (*Dependency, error) {
			return &Dependency{up}, nil
		},
	)
}

func ExampleWith2() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency1 struct{}
	type UpstreamDependency2 struct{}
	type Dependency struct {
		Up1 *UpstreamDependency1
		Up2 *UpstreamDependency2
	}

	// Declare a constructor for the Dependency type. It depends on the two
	// upstream dependency types (which are assumed to be declared elsewhere).
	imbue.With2(
		con,
		func(
			ctx *imbue.Context,
			up1 *UpstreamDependency1,
			up2 *UpstreamDependency2,
		) (*Dependency, error) {
			return &Dependency{up1, up2}, nil
		},
	)
}
