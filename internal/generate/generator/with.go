package generator

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// generateWithFunc generates the WithX function for the given number of
// dependencies.
func generateWithFunc(code *jen.File, depCount int) {
	name := fmt.Sprintf("With%d", depCount)

	switch depCount {
	case 0:
		code.Commentf(
			"%s describes how to construct values of type T.",
			name,
		)
	case 1:
		code.Commentf(
			"%s describes how to construct values of type T from a single dependency.",
			name,
		)
	default:
		code.Commentf(
			"%s describes how to construct values of type T from %d dependencies.",
			name,
			depCount,
		)
	}

	code.
		Func().
		Id(name).
		Types(
			types(true, depCount)...,
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
			jen.Line(),
		).
		BlockFunc(func(g *jen.Group) {
			generateWithFuncBody(depCount, g)
		})
}

func generateWithFuncBody(depCount int, code *jen.Group) {
	code.
		List(
			jen.Id("_"),
			jen.Id("file"),
			jen.Id("line"),
			jen.Id("_"),
		).
		Op(":=").
		Qual("runtime", "Caller").
		Call(
			jen.Lit(1),
		)

	code.Line()

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
				Add(declaringDeclVar(depCount)).Dot("Declare").
				Call(
					jen.Line().
						Id("file"),
					jen.Line().
						Id("line"),
					jen.Line().
						Func().
						Params().
						Params(
							jen.
								Id("constructor").
								Types(
									declaringType(depCount),
								),
							jen.Error(),
						).
						BlockFunc(func(g *jen.Group) {
							generateDeclareFuncBody(depCount, g)
						}),
					jen.Line(),
				).
				Op(";").
				Err().Op("!=").Nil(),
		).
		Block(
			jen.Panic(jen.Err()),
		)
}

func generateDeclareFuncBody(depCount int, code *jen.Group) {
	if depCount == 0 {
		code.
			Return(
				jen.Id("fn"),
				jen.Nil(),
			)

		return
	}

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
					Qual(pkgPath, "addDependency").
					Call(
						declaringDeclVar(depCount),
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
				).
				Params(
					declaringVar(depCount).Add(declaringType(depCount)),
					jen.Id("_").Error(),
				).
				BlockFunc(func(g *jen.Group) {
					generateConstructorFuncBody(depCount, g)
				}),
			jen.
				Nil(),
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
				Id("fn").
				Call(
					inputVars(depCount)...,
				),
		)
}
