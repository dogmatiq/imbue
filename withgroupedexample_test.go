package imbue_test

import (
	"context"
	"fmt"

	"github.com/dogmatiq/imbue"
)

// Connection is a connection to a remote service.
type Connection string

// Client is an API client that uses a connection to a remote service.
type Client struct {
	Conn Connection
}

type (
	// ServiceA is a group for the dependencies related to service A.
	ServiceA imbue.Group

	// ServiceB is a group for the dependencies related to service B.
	ServiceB imbue.Group
)

func Example_groupedDependencies() {
	con := imbue.New()
	defer con.Close()

	// Declare a constructor for the Connection to ServiceA.
	imbue.With0Grouped[ServiceA](
		con,
		func(
			ctx *imbue.Context,
		) (Connection, error) {
			return "<connection-a>", nil
		},
	)

	// Declare a constructor for the API Client for ServiceA.
	imbue.With1Grouped[ServiceA](
		con,
		func(
			ctx *imbue.Context,
			conn imbue.FromGroup[ServiceA, Connection],
		) (Client, error) {
			return Client{conn.Value}, nil
		},
	)

	// Declare a constructor for the Connection to ServiceB.
	imbue.With0Grouped[ServiceB](
		con,
		func(
			ctx *imbue.Context,
		) (Connection, error) {
			return "<connection-b>", nil
		},
	)

	// Declare a constructor for the API Client for ServiceB.
	imbue.With1Grouped[ServiceB](
		con,
		func(
			ctx *imbue.Context,
			conn imbue.FromGroup[ServiceB, Connection],
		) (Client, error) {
			return Client{conn.Value}, nil
		},
	)

	// Invoke a function that depends on both the Foreground and Background
	// colors.
	err := imbue.Invoke2(
		context.Background(),
		con,
		func(
			ctx context.Context,
			clientA imbue.FromGroup[ServiceA, Client],
			clientB imbue.FromGroup[ServiceB, Client],
		) error {
			// Grouped dependencies have a Value field which in this case is the
			// Client value.
			fmt.Println("service A:", clientA.Value)
			fmt.Println("service B:", clientB.Value)
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	// Output:
	// service A: {<connection-a>}
	// service B: {<connection-b>}
}
