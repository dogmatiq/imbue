package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// generateInvokeFunc generates the InvokeX function for the given number of
// dependencies.
func generateInvokeFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Invoke%d", depCount)

	switch depCount {
	case 1:
		code.
			Commentf(
				"%s calls a function with a single dependency.",
				name,
			)
	default:
		code.
			Commentf(
				"%s calls a function with %d dependencies.",
				name,
				depCount,
			)
	}

	impl := &jen.Statement{}

	impl.Add(
		jen.Id("rctx").Op(":=").Qual(pkgPath, "rootContext").
			Call(
				contextName(),
			),
	)

	impl.Add(
		jen.Line(),
	)

	for n := 0; n < depCount; n++ {
		impl.Add(
			jen.List(
				jen.Id(dependencyParamName(depCount, n)),
				jen.Err(),
			).
				Op(":=").
				Qual(pkgPath, "resolve").
				Index(
					jen.Id(dependencyTypeName(depCount, n)),
				).
				Call(
					jen.Id("rctx"),
					containerName(),
				),
		)

		impl.Add(
			jen.If(
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(
					jen.Err(),
				),
			),
		)

		impl.Add(
			jen.Line(),
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
		Types(typeParams(false, depCount)...).
		Params(
			jen.Line().
				Add(stdContextParam()),
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputParamTypes(stdContextType(), depCount)...,
				).
				Params(
					jen.Error(),
				),
			jen.Line(),
		).
		Params(
			jen.Error(),
		).
		Block(*impl...)
}
