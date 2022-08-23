package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateGo generates the GoX() functions.
func GenerateGo(code *jen.File) {
	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateGoFunc(code, depCount)
	}
}

// generateGoFunc generates the GoX function for the given number of
// dependencies.
func generateGoFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Go%d", depCount)

	switch depCount {
	case 1:
		code.Commentf("%s starts a new goroutine by calling a function with a single dependency.", name)
	default:
		code.Commentf("%s starts a new goroutine by calling a function with %d dependencies.", name, depCount)
	}

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
									jen.
										Qual(pkgPath, fmt.Sprintf("Invoke%d", depCount)).
										Call(
											waitGroupVar().Dot("ctx"),
											waitGroupVar().Dot("con"),
											invokeFuncVar(),
											jen.Id("options").Op("..."),
										),
								),
						),
				),
		)
}
