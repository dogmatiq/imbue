package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateWithNamed generates the WithXNamed() functions.
func GenerateWithNamed(code *jen.File) {
	for depCount := 0; depCount <= maxDependencies; depCount++ {
		generateWithNamedFunc(code, depCount)
	}
}

// generateWithNamedFunc generates the WithXNamed function for the given number
// of dependencies.
func generateWithNamedFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("With%dNamed", depCount)

	switch depCount {
	case 0:
		code.Commentf("%s describes how to construct named values of type T.", name)
	case 1:
		code.Commentf("%s describes how to construct named values of type T from a single", name)
		code.Commentf("dependency.")
	default:
		code.Commentf("%s describes how to construct named values of type T from %d", name, depCount)
		code.Commentf("dependencies.")
	}

	code.Comment("")
	code.Commentf("%s is the name given to the dependency.", namedTypeString(depCount))
	code.Commentf("%s is the type of the dependency.", declaringTypeString(depCount))

	code.
		Func().
		Id(name).
		Types(
			types(
				depCount,
				namedType(depCount).Qual(pkgPath, "Name").
					Types(
						declaringType(depCount),
					),
				declaringType(depCount),
			)...,
		).
		Params(
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Id("fn").
				Func().
				Params(
					inputTypes(imbueContextType(), depCount)...,
				).
				Params(
					declaringType(depCount),
					jen.Error(),
				),
			jen.Line().
				Id("options").
				Op("...").
				Qual(pkgPath, "WithNamedOption"),
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateWithNamedFuncBody(depCount, g)
		})
}

func generateWithNamedFuncBody(depCount int, code *jen.Group) {
	code.
		Id(fmt.Sprintf("With%d", depCount)).
		Call(
			jen.Line().
				Add(containerVar()),
			jen.Line().
				Func().
				Params(
					inputParams(imbueContextType(), depCount)...,
				).
				Params(
					jen.Qual(pkgPath, "ByName").
						Types(
							namedType(depCount),
							declaringType(depCount),
						),
					jen.Error(),
				).
				Block(
					jen.
						List(
							declaringVar(depCount),
							jen.Err(),
						).
						Op(":=").
						Id("fn").
						Call(
							inputVars(depCount)...,
						),
					jen.
						Return(
							jen.Qual(pkgPath, "withName").
								Types(
									namedType(depCount),
								).
								Call(
									declaringVar(depCount),
								),
							jen.Err(),
						),
				),
			jen.Line(),
		)
}
