package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateWith generates the WithX() functions.
func GenerateWith(code *jen.File) {
	for depCount := 0; depCount <= maxDependencies; depCount++ {
		generateWithFunc(code, depCount)
	}
}

// generateWithFunc generates the WithX function for the given number of
// dependencies.
func generateWithFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("With%d", depCount)

	switch depCount {
	case 0:
		code.Commentf("%s describes how to construct values of type T.", name)
	case 1:
		code.Commentf("%s describes how to construct values of type T from a single dependency.", name)
	default:
		code.Commentf("%s describes how to construct values of type T from %d dependencies.", name, depCount)
	}

	code.
		Func().
		Id(name).
		Types(
			types(
				depCount,
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
				Qual(pkgPath, "WithOption"),
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateWithFuncBody(depCount, g)
		})
}

func generateWithFuncBody(depCount int, code *jen.Group) {
	code.
		Add(declaringDeclVar(depCount)).
		Op(":=").
		Qual(pkgPath, "get").
		Types(
			declaringType(depCount),
		).
		Call(
			containerVar(),
		)

	for n := 0; n < depCount; n++ {
		code.
			Add(dependencyDeclVar(depCount, n)).
			Op(":=").
			Qual(pkgPath, "get").
			Types(
				dependencyType(depCount, n),
			).
			Call(
				containerVar(),
			)
	}

	code.Line()

	code.
		Add(declaringDeclVar(depCount)).Dot("Declare").
		CallFunc(
			func(code *jen.Group) {
				code.
					Line().
					Func().
					Params(
						imbueContextParam(),
					).
					Params(
						declaringVar(depCount).Add(declaringType(depCount)),
						jen.Id("_").Error(),
					).
					BlockFunc(func(g *jen.Group) {
						generateConstructorFuncBody(depCount, g)
					})

				for n := 0; n < depCount; n++ {
					code.
						Line().
						Add(dependencyDeclVar(depCount, n))
				}

				code.Line()
			},
		)
}

func generateConstructorFuncBody(depCount int, code *jen.Group) {
	for n := 0; n < depCount; n++ {
		code.
			List(
				dependencyVar(depCount, n),
				jen.Err(),
			).
			Op(":=").
			Add(dependencyDeclVar(depCount, n)).Dot("Resolve").
			Call(
				contextVar(),
			)

		code.
			If(
				jen.Err().Op("!=").Nil(),
			).
			Block(
				jen.
					Return(
						declaringVar(depCount),
						jen.Err(),
					),
			)

		code.Line()
	}

	code.
		Return(
			jen.
				Add(constructorVar()).
				Call(
					inputVars(depCount, contextVar())...,
				),
		)
}
