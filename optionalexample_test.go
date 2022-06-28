package imbue_test

// func _ExampleOptional() {
// 	con := imbue.New()
// 	defer con.Close()

// 	// Declare a type to use as a dependency within the example.
// 	type Dependency struct{}

// 	// Invoke a function that optionally depends on the Dependency type.
// 	if err := imbue.Invoke1(
// 		context.Background(),
// 		con,
// 		func(
// 			ctx context.Context,
// 			dep imbue.Optional[Dependency],
// 		) error {
// 			if dep.Ok() {
// 				fmt.Println("dependency is available: ", dep.Value())
// 			} else {
// 				fmt.Println("dependency is unavailable")
// 			}

// 			return nil
// 		},
// 	); err != nil {
// 		panic(err)
// 	}

// 	// Print the dependency tree.
// 	fmt.Println(con)

// 	// Output:
// 	// <container>
// 	// └── *imbue_test.Dependency
// }
