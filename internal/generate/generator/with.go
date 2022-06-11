package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// generateWithFunc generates the WithX function for the given number of
// dependencies.
func generateWithFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("With%d", depCount)

	switch depCount {
	case 0:
		code.
			Commentf(
				"%s describes how to construct values of type T.",
				name,
			)
	case 1:
		code.
			Commentf(
				"%s describes how to construct values of type T from a single dependency.",
				name,
			)
	default:
		code.
			Commentf(
				"%s describes how to construct values of type T from %d dependencies.",
				name,
				depCount,
			)
	}

	impl := &jen.Statement{}

	for n := 0; n < depCount; n++ {
		impl.Add(
			jen.List(
				jen.Id(dependencyParamName(depCount, n)),
				jen.Err(),
			).
				Op(":=").
				Qual(pkgPath, "get").
				Index(
					jen.Id(dependencyTypeName(depCount, n)),
				).
				Call(
					containerName(),
				),
		)

		impl.Add(
			jen.If(
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(
					jen.Id("value"),
					jen.Err(),
				),
			),
		)
	}

	impl.Add(
		jen.Return(
			jen.Id("fn").Call(
				inputParamNames(depCount)...,
			),
		),
	)

	code.
		Func().
		Id(name).
		Types(typeParams(true, depCount)...).
		Params(
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputParamTypes(imbueContextType(), depCount)...,
				).
				Params(
					jen.Id("T"),
					jen.Error(),
				),
			jen.Line(),
		).
		Block(
			jen.Qual(pkgPath, "register").
				Call(
					jen.Line().
						Add(containerName()),
					jen.Line().
						Func().
						Params(
							imbueContextParam(),
							containerParam(),
						).
						Params(
							jen.Id("value").Id("T"),
							jen.Id("_").Error(),
						).
						Block(*impl...),
					jen.Line(),
				),
		)
}
