package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// GenerateInject generates the InjectX() functions.
func GenerateInject(code *jen.File) {
	for depCount := 1; depCount <= maxDependencies; depCount++ {
		generateInjectFunc(code, depCount)
	}
}

// generateInjectFunc generates the InjectX function for the given number of
// dependencies.
func generateInjectFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("Inject%d", depCount)

	switch depCount {
	case 1:
		code.Commentf("%s describes how to initialize values of type T after it is constructed", name)
		code.Commentf("using a single additional dependency.")
	default:
		code.Commentf("%s describes how to initialize values of type T after it is constructed", name)
		code.Commentf("using additional %d dependencies.", depCount)
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
				Qual(pkgPath, "InjectOption"),
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateInjectFuncBody(depCount, g)
		})
}

func generateInjectFuncBody(depCount int, code *jen.Group) {
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
					Dot("AddDependency").
					Call(
						dependencyDeclVar(depCount, n),
					).
					Op(";").
					Err().Op("!=").Nil(),
			).
			Block(
				jen.Panic(
					jen.Err(),
				),
			)

		code.Line()
	}

	code.Line()

	code.
		Add(declaringDeclVar(depCount)).Dot("AddInitializer").
		Call(
			jen.Line().
				Func().
				Params(
					imbueContextParam(),
					declaringVar(depCount).Add(declaringType(depCount)),
				).
				Params(
					jen.Error(),
				).
				BlockFunc(func(g *jen.Group) {
					generateInitFuncBody(depCount, g)
				}),
			jen.Line(),
		)
}

func generateInitFuncBody(depCount int, code *jen.Group) {
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
