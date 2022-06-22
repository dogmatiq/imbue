package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
	"github.com/dogmatiq/imbue/internal/generate/generator"
)

func main() {
	if err := generator.Generate(
		os.Args[2],
		getGenerator(os.Args[2]),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getGenerator(filename string) func(*jen.File) {
	switch filepath.Base(filename) {
	case "decorate.gen.go":
		return generator.GenerateDecorate
	case "invoke.gen.go":
		return generator.GenerateInvoke
	case "with.gen.go":
		return generator.GenerateWith
	case "withgrouped.gen.go":
		return generator.GenerateWithGrouped
	case "withnamed.gen.go":
		return generator.GenerateWithNamed
	}

	panic("could not determine generator for " + filename)
}
