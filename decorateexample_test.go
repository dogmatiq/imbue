package imbue_test

import (
	"fmt"
	"net/http"

	"github.com/dogmatiq/imbue"
)

func ExampleDecorate0() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type DecoratedDependency struct {
		Value string
	}

	// Declare a decorator for the DecoratedDependency type.
	imbue.Decorate0(
		con,
		func(
			ctx *imbue.Context,
			d *DecoratedDependency,
		) (*DecoratedDependency, error) {
			d.Value = "<decorated value>"
			return d, nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *imbue_test.DecoratedDependency
}

func ExampleDecorate1() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency struct{}
	type DecoratedDependency struct {
		Up *UpstreamDependency
	}

	// Declare a decorator for the DecoratedDependency type. It depends on the
	// upstream dependency type (which is assumed to be declared elsewhere).
	imbue.Decorate1(
		con,
		func(
			ctx *imbue.Context,
			d *DecoratedDependency,
			up *UpstreamDependency,
		) (*DecoratedDependency, error) {
			d.Up = up
			return d, nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *imbue_test.DecoratedDependency
	//     └── *imbue_test.UpstreamDependency
}

func ExampleDecorate2() {
	con := imbue.New()
	defer con.Close()

	// Declare some types to use as dependencies within the example.
	type UpstreamDependency1 struct{}
	type UpstreamDependency2 struct{}
	type DecoratedDependency struct {
		Up1 *UpstreamDependency1
		Up2 *UpstreamDependency2
	}

	// Declare a decorator for the DecoratedDependency type. It depends on the
	// two upstream dependency types (which are assumed to be declared
	// elsewhere).
	imbue.Decorate2(
		con,
		func(
			ctx *imbue.Context,
			d *DecoratedDependency,
			up1 *UpstreamDependency1,
			up2 *UpstreamDependency2,
		) (*DecoratedDependency, error) {
			d.Up1 = up1
			d.Up2 = up2
			return d, nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *imbue_test.DecoratedDependency
	//     ├── *imbue_test.UpstreamDependency1
	//     └── *imbue_test.UpstreamDependency2
}

func ExampleDecorate0_httpServeMux() {
	con := imbue.New()
	defer con.Close()

	// This example illustrates how decoration can be used to add routes to an
	// http.ServeMux that is declared in a separate location.

	imbue.With0(
		con,
		func(
			ctx *imbue.Context,
		) (*http.ServeMux, error) {
			return http.NewServeMux(), nil
		},
	)

	imbue.Decorate0(
		con,
		func(
			ctx *imbue.Context,
			mux *http.ServeMux,
		) (*http.ServeMux, error) {
			mux.HandleFunc("/account", accountHandler)
			mux.HandleFunc("/dashboard", dashboardHandler)
			return mux, nil
		},
	)

	// Print the dependency tree.
	fmt.Println(con)

	// Output:
	// <container>
	// └── *http.ServeMux
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "account handler")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dashboard handler")
}
