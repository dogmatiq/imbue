package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// generateCallWithFunc generates the CallWithX function for the given number of
// dependencies.
func generateCallWithFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("CallWith%d", depCount)

	switch depCount {
	case 1:
		code.
			Commentf(
				"%s calls a function that returns a value of type T with a single dependency.",
				name,
			)
	default:
		code.
			Commentf(
				"%s calls a functionthat returns a value of type T with %d dependencies.",
				name,
				depCount,
			)
	}

	code.
		Func().
		Id(name).
		Types(typeParams(true, depCount)...).
		Params(
			jen.Line().
				Add(stdContextParam()),
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputParamNames(depCount)...,
				).
				Params(
					jen.Id("T"),
					jen.Error(),
				),
			jen.Line(),
		).
		Params(
			jen.Id("T"),
			jen.Error(),
		).
		Block(
			jen.Var().Id("result").Id("T"),
			jen.Line(),
			jen.Err().Op(":=").
				Id(fmt.Sprintf("InvokeWith%d", depCount)).
				Call(
					jen.Line().
						Add(contextName()),
					jen.Line().
						Add(containerName()),
					jen.Line().
						Func().
						Params(
							inputParams(stdContextType(), depCount)...,
						).
						Params(
							jen.Error(),
						).
						Block(
							jen.Var().Err().Error(),
							jen.List(
								jen.Id("result"),
								jen.Err(),
							).Op("=").Id("fn").Call(
								inputParamNames(depCount)...,
							),
							jen.Return(
								jen.Err(),
							),
						),
					jen.Line(),
				),
			jen.Line(),
			jen.Return(
				jen.Id("result"),
				jen.Err(),
			),
		)
}
