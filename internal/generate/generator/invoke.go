package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateInvoke generates the InvokeX() methods on Container and the
// deprecated InvokeX() functions that forward to them.
func GenerateInvoke(code *jen.File) {
	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateInvokeMethod(code, depCount)
	}

	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateInvokeFunc(code, depCount)
	}
}

// invokeComment returns the descriptive part of the InvokeX documentation
// comment for the given number of dependencies.
func invokeComment(depCount int) string {
	if depCount == 1 {
		return "calls a function with a single dependency."
	}

	return fmt.Sprintf("calls a function with %d dependencies.", depCount)
}

// generateInvokeMethod generates the InvokeX method on Container for the
// given number of dependencies.
func generateInvokeMethod(code *jen.File, depCount int) {
	name := fmt.Sprintf("Invoke%d", depCount)

	code.Commentf("%s %s", name, invokeComment(depCount))

	code.
		Func().
		Params(
			containerParam(),
		).
		Id(name).
		Types(
			types(depCount)...,
		).
		Params(
			jen.Line().
				Add(stdContextParam()),
			jen.Line().
				Add(invokeFuncVar()).
				Func().
				Params(
					inputTypes(depCount, stdContextType())...,
				).
				Params(
					jen.Error(),
				),
			jen.Line().
				Id("options").
				Op("...").
				Qual(pkgPath, "InvokeOption"),
			jen.Line(),
		).
		Params(
			jen.Error(),
		).
		BlockFunc(func(g *jen.Group) {
			generateInvokeMethodBody(depCount, g)
		})
}

func generateInvokeMethodBody(depCount int, code *jen.Group) {
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
				jen.Id("ctx"),
			)

		code.
			If(
				jen.Err().Op("!=").Nil(),
			).
			Block(
				jen.Return(
					jen.
						Qual(pkgPath, "filterInvokeError").
						Call(
							jen.Err(),
						),
				),
			)

		code.Line()
	}

	code.Return(
		jen.
			Qual(pkgPath, "filterInvokeError").
			Call(
				jen.
					Add(invokeFuncVar()).
					Call(
						inputVars(depCount, contextVar())...,
					),
			),
	)
}

// generateInvokeFunc generates the deprecated InvokeX function that forwards
// to the InvokeX method for the given number of dependencies.
func generateInvokeFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Invoke%d", depCount)

	code.Commentf("%s %s", name, invokeComment(depCount))
	code.Comment("//")
	code.Commentf("Deprecated: Use [Container.%s] instead.", name)
	code.Comment("//")
	code.Comment("//go:fix inline")

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
				Add(invokeFuncVar()).
				Func().
				Params(
					inputTypes(depCount, stdContextType())...,
				).
				Params(
					jen.Error(),
				),
			jen.Line().
				Id("options").
				Op("...").
				Qual(pkgPath, "InvokeOption"),
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
	code.Return(
		containerVar().
			Dot(fmt.Sprintf("Invoke%d", depCount)).
			Call(
				contextVar(),
				invokeFuncVar(),
				jen.Id("options").Op("..."),
			),
	)
}
