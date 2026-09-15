package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateGo generates the GoX() methods on WaitGroup and the deprecated GoX()
// functions that forward to them.
func GenerateGo(code *jen.File) {
	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateGoMethod(code, depCount)
	}

	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateGoFunc(code, depCount)
	}
}

// goComment returns the descriptive part of the GoX documentation comment for
// the given number of dependencies.
func goComment(depCount int) string {
	if depCount == 1 {
		return "starts a new goroutine by calling a function with a single dependency."
	}

	return fmt.Sprintf("starts a new goroutine by calling a function with %d dependencies.", depCount)
}

// generateGoMethod generates the GoX method on WaitGroup for the given number
// of dependencies.
func generateGoMethod(code *jen.File, depCount int) {
	name := fmt.Sprintf("Go%d", depCount)

	code.Commentf("%s %s", name, goComment(depCount))

	code.
		Func().
		Params(
			waitGroupParam(),
		).
		Id(name).
		Types(
			types(depCount)...,
		).
		Params(
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
		BlockFunc(func(g *jen.Group) {
			generateGoMethodBody(depCount, g)
		})
}

func generateGoMethodBody(depCount int, code *jen.Group) {
	code.
		Add(
			waitGroupVar().
				Dot("group").
				Dot("Go").
				Call(
					jen.
						Func().
						Params().
						Error().
						Block(
							jen.
								Return(
									waitGroupVar().
										Dot("con").
										Dot(fmt.Sprintf("Invoke%d", depCount)).
										Call(
											waitGroupVar().Dot("ctx"),
											invokeFuncVar(),
											jen.Id("options").Op("..."),
										),
								),
						),
				),
		)
}

// generateGoFunc generates the deprecated GoX function that forwards to the GoX
// method for the given number of dependencies.
func generateGoFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Go%d", depCount)

	code.Commentf("%s %s", name, goComment(depCount))
	code.Comment("//")
	code.Commentf("Deprecated: Use [WaitGroup.%s] instead.", name)
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
				Add(waitGroupParam()),
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
		BlockFunc(func(g *jen.Group) {
			generateGoFuncBody(depCount, g)
		})
}

func generateGoFuncBody(depCount int, code *jen.Group) {
	code.
		Add(
			waitGroupVar().
				Dot(fmt.Sprintf("Go%d", depCount)).
				Call(
					invokeFuncVar(),
					jen.Id("options").Op("..."),
				),
		)
}
