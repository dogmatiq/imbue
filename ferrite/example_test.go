package ferrite_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/imbue/ferrite"
)

var (
	// APIHostName declares the API_HOST_NAME environment variable as a string.
	APIHostName = ferrite.String("API_HOST_NAME")

	// APIPort declares the API_PORT environment variable as a uint16.
	APIPort = ferrite.Unsigned[uint16]("API_PORT")
)

func Example() {
	// Set some environment variables to use in our example.
	os.Setenv("API_HOST_NAME", "server.example.org")
	os.Setenv("API_PORT", "8080")

	// Resolve the values of our environment variables.
	fmt.Printf("%s = %v\n", APIHostName.Name(), APIHostName.Value())
	fmt.Printf("%s = %v\n", APIPort.Name(), APIPort.Value())

	// Output:
	// API_HOST_NAME = server.example.org
	// API_PORT = 8080
}
