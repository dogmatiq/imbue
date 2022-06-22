package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateWithGrouped generates the WithXGrouped() functions.
func GenerateWithGrouped(code *jen.File) {
	for depCount := 0; depCount <= maxDependencies; depCount++ {
		generateWithGroupedFunc(code, depCount)
	}
}

// generateWithGroupedFunc generates the WithXGrouped function for the given
// number of dependencies.
func generateWithGroupedFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("With%dGrouped", depCount)

	switch depCount {
	case 0:
		code.Commentf("%s describes how to construct grouped values of type T.", name)
	case 1:
		code.Commentf("%s describes how to construct grouped values of type T from a", name)
		code.Commentf("single dependency.")
	default:
		code.Commentf("%s describes how to construct grouped values of type T from %d", name, depCount)
		code.Commentf("dependencies.")
	}

	code.Comment("")
	code.Commentf("%s is the group that contains the dependency.", groupedTypeString(depCount))
	code.Commentf("%s is the type of the dependency.", declaringTypeString(depCount))

	code.
		Func().
		Id(name).
		Types(
			types(
				depCount,
				groupedType(depCount).Qual(pkgPath, "Group"),
				declaringType(depCount),
			)...,
		).
		Params(
			jen.Line().
				Add(containerParam()),
			jen.Line().
				Add(constructorVar()).
				Func().
				Params(
					inputTypes(depCount, imbueContextType())...,
				).
				Params(
					declaringType(depCount),
					jen.Error(),
				),
			jen.Line().
				Id("options").
				Op("...").
				Qual(pkgPath, "WithGroupedOption"),
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateWithGroupedFuncBody(depCount, g)
		})
}

func generateWithGroupedFuncBody(depCount int, code *jen.Group) {
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
					jen.Qual(pkgPath, "FromGroup").
						Types(
							groupedType(depCount),
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
						Add(constructorVar()).
						Call(
							inputVars(depCount, contextVar())...,
						),
					jen.
						Return(
							jen.Qual(pkgPath, "inGroup").
								Types(
									groupedType(depCount),
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
