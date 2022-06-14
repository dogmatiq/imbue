package imbue

// Name is a type/constraint for declaring named dependencies.
type Name[T any] interface {
	nameOf(T)
}

// ByName declares a dependency on a named type.
type ByName[N Name[T], T any] struct {
	// Value is the dependency itself.
	Value T
}

// withName wraps a value of type T to present it as a ByName[N, T].
func withName[N Name[T], T any](v T) ByName[N, T] {
	return ByName[N, T]{
		Value: v,
	}
}
