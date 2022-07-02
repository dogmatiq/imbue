package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateDecorate generates the DecorateX() functions.
func GenerateDecorate(code *jen.File) {
	for depCount := 0; depCount <= maxDependencies; depCount++ {
		generateDecorateFunc(code, depCount)
	}
}

// generateDecorateFunc generates the DecorateX function for the given number of
// dependencies.
func generateDecorateFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Decorate%d", depCount)

	switch depCount {
	case 0:
		code.Commentf("%s describes how to decorate values of type T after construction.", name)
	case 1:
		code.Commentf("%s describes how to decorate values of type T after construction using", name)
		code.Commentf("a single additional dependency.")
	default:
		code.Commentf("%s describes how to decorate values of type T after construction using", name)
		code.Commentf("%d additional dependencies.", depCount)
	}

	code.Commentf("")
	code.Commentf("The dependency being decorated is passed to %s and replaced with", decoratorFuncName)
	code.Commentf("the decorator's return value.")
	code.Commentf("")
	code.Commentf("The decorated dependency may be manipulated in-place.")

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
				Add(decoratorVar()).
				Func().
				Params(
					inputTypes(
						depCount,
						imbueContextType(),
						declaringType(depCount),
					)...,
				).
				Params(
					declaringType(depCount),
					jen.Error(),
				),
			jen.Line().
				Id("options").
				Op("...").
				Qual(pkgPath, "DecorateOption"),
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateDecorateFuncBody(depCount, g)
		})
}

func generateDecorateFuncBody(depCount int, code *jen.Group) {
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
		Add(declaringDeclVar(depCount)).Dot("Decorate").
		CallFunc(
			func(code *jen.Group) {
				code.
					Line().
					Func().
					Params(
						imbueContextParam(),
						declaringVar(depCount).Add(declaringType(depCount)),
					).
					Params(
						declaringType(depCount),
						jen.Error(),
					).
					BlockFunc(func(g *jen.Group) {
						generateDecoratorFuncBody(depCount, g)
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

func generateDecoratorFuncBody(depCount int, code *jen.Group) {
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
				Add(decoratorVar()).
				Call(
					inputVars(
						depCount,
						contextVar(),
						declaringVar(depCount),
					)...,
				),
		)
}
