package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/dogmatiq/imbue/internal/generate/generator"
)

func main() {
	buf := &bytes.Buffer{}
	if err := generator.Generate(buf); err != nil {
		fmt.Fprintf(os.Stderr, "unable to generate code: %s\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(
		os.Args[2],
		buf.Bytes(),
		0644,
	); err != nil {
		fmt.Fprintf(os.Stderr, "unable to write to file: %s\n", err)
		os.Exit(2)
	}
}
