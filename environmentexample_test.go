package imbue_test

import (
	"context"
	"fmt"
	"os"

	"github.com/dogmatiq/imbue"
)

type (
	// API_HOST_NAME declares the API_HOST_NAME environment variable as a string.
	API_HOST_NAME imbue.EnvironmentVariable[string]

	// API_PORT declares the API_PORT environment variable as a uint16.
	API_PORT imbue.EnvironmentVariable[uint16]
)

func Example_environmentVariables() {
	// Set some environment variables to use in our example.
	os.Setenv("API_HOST_NAME", "server.example.org")
	os.Setenv("API_PORT", "8080")

	con := imbue.New()
	defer con.Close()

	// Invoke a function that depends on both the API_HOST_NAME and API_PORT
	// environment variables.
	err := imbue.Invoke2(
		context.Background(),
		con,
		func(
			ctx context.Context,
			h imbue.FromEnvironment[API_HOST_NAME, string],
			p imbue.FromEnvironment[API_PORT, uint16],
		) error {
			fmt.Println(h.Name(), "=", h.Value())
			fmt.Println(p.Name(), "=", p.Value())
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	// Output:
	// API_HOST_NAME = server.example.org
	// API_PORT = 8080
}
