package imbue

// Name is a constraint for a type that identifies a named dependency.
//
// Names are used to distinguish between multiple dependencies of the same type.
//
// Names are declared by declaring a type that uses imbue.Name[T] as its
// underlying type, where T is the type of the dependency being named.
type Name[T any] interface {
	nameOf(T)
}

// ByName requests a dependency on a type with a specific name.
//
// It is used as a parameter type to user-defined functions passed to WithX(),
// DecorateX() and InvokeX() to request a dependency of type T that is named N.
type ByName[N Name[T], T any] struct {
	value T
}

// Name returns the name given to the dependency.
func (v ByName[N, T]) Name() string {
	return typeOf[N]().Name()
}

// Value returns the dependency value.
func (v ByName[N, T]) Value() T {
	return v.value
}

// withName wraps a value of type T to present it as a ByName[N, T].
func withName[N Name[T], T any](v T) ByName[N, T] {
	return ByName[N, T]{
		value: v,
	}
}

// WithNamedOption is an option that changes the behavior of a call to
// WithXNamed().
type WithNamedOption interface {
	applyWithNamedOption()
}
