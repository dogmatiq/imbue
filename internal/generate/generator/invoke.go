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
		code.Commentf(
			"%s calls a function with a single dependency.",
			name,
		)
	default:
		code.Commentf(
			"%s calls a function with %d dependencies.",
			name,
			depCount,
		)
	}

	code.
		Func().
		Id(name).
		Types(
			types(depCount)...,
		).
		Params(
			jen.Line().
				Add(stdContextParam()),
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputTypes(stdContextType(), depCount)...,
				).
				Params(
					jen.Error(),
				),
			jen.Line(),
		).
		Params(
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			generateInvokeFuncBody(depCount, g)
		})
}

func generateInvokeFuncBody(depCount int, code *jen.Group) {
	code.
		Id("rctx").
		Op(":=").
		Qual(pkgPath, "rootContext").
		Call(
			contextVar(),
			containerVar(),
		)

	code.Line()

	for n := 0; n < depCount; n++ {
		code.
			List(
				dependencyVar(depCount, n),
				jen.Err(),
			).
			Op(":=").
			Qual(pkgPath, "get").
			Types(
				dependencyType(depCount, n),
			).
			Call(
				containerVar(),
			).
			Dot("Resolve").
			Call(
				jen.Id("rctx"),
			)

		code.If(
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(
				jen.Err(),
			),
		)

		code.Line()
	}

	code.Return(
		jen.
			Id("fn").
			Call(
				inputVars(depCount)...,
			),
	)
}
