package imbue

// Optional represents an optional dependency of type T.
type Optional[T any] struct {
	value T
	err   error
}

// Value returns the dependency value if it is available; otherwise, it returns
// a non-nil error.
//
// A dependency is considered unavailable if it does not have a constructor
// declared, or if that constructor returns an error.
func (v Optional[T]) Value() (T, error) {
	return v.value, v.err
}

func (Optional[T]) declare(
	con *Container,
	decl *declarationOf[Optional[T]],
) {
	dep := get[T](con)

	decl.Declare(
		func(ctx Context) (Optional[T], error) {
			v, err := dep.Resolve(ctx)
			return Optional[T]{v, err}, nil
		},
		dep,
	)
}
