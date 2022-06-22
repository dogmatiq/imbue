package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateDecorate generates the DecorateX() functions.
func GenerateDecorate(code *jen.File) {
	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateDecorateFunc(code, depCount)
	}
}

// generateDecorateFunc generates the DecorateX function for the given number of
// dependencies.
func generateDecorateFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Decorate%d", depCount)

	switch depCount {
	case 1:
		code.Commentf("%s describes how to decorate values of type T after construction using", name)
		code.Commentf("a single additional dependency.")
	default:
		code.Commentf("%s describes how to decorate values of type T after construction using", name)
		code.Commentf("%d additional dependencies.", depCount)
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
				Id("fn").
				Func().
				Params(
					inputTypes(
						depCount,
						imbueContextType(),
						declaringType(depCount),
					)...,
				).
				Params(
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

	code.Line()

	code.
		If(
			jen.Err().
				Op(":=").
				Add(declaringDeclVar(depCount)).Dot("AddDecorator").
				Call(
					jen.Line().
						Func().
						Params().
						Params(
							jen.
								Id("decorator").
								Types(
									declaringType(depCount),
								),
							jen.Error(),
						).
						BlockFunc(func(g *jen.Group) {
							generateDecoratorFactoryFuncBody(depCount, g)
						}),
					jen.Line(),
				).
				Op(";").
				Err().Op("!=").Nil(),
		).
		Block(
			jen.Panic(
				jen.Err(),
			),
		)
}

func generateDecoratorFactoryFuncBody(depCount int, code *jen.Group) {
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

		code.
			If(
				jen.
					Err().
					Op(":=").
					Add(declaringDeclVar(depCount)).
					Dot("AddDecoratorDependency").
					Call(
						dependencyDeclVar(depCount, n),
					).
					Op(";").
					Err().Op("!=").Nil(),
			).
			Block(
				jen.Return(
					jen.Nil(),
					jen.Err(),
				),
			)

		code.Line()
	}

	code.
		Return(
			jen.
				Func().
				Params(
					imbueContextParam(),
					declaringVar(depCount).Add(declaringType(depCount)),
				).
				Params(
					jen.Error(),
				).
				BlockFunc(func(g *jen.Group) {
					generateDecoratorFuncBody(depCount, g)
				}),
			jen.
				Nil(),
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
						jen.Err(),
					),
			)

		code.Line()
	}

	code.
		Return(
			jen.
				Id("fn").
				Call(
					inputVars(
						depCount,
						contextVar(),
						declaringVar(depCount),
					)...,
				),
		)
}
