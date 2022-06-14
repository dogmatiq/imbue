package imbue

// Name is a type/constraint for declaring named dependencies.
type Name[T any] interface {
	nameOf(T)
}

// ByName declares a dependency on a named type.
//
// It is a function that returns the dependency of type T that is named N.
type ByName[N Name[T], T any] func() T

// withName wraps a value of type T to present it as a ByName[N, T] function.
func withName[N Name[T], T any](v T) ByName[N, T] {
	return func() T {
		return v
	}
}
