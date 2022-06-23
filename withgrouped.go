package imbue

// Group is a constraint for a type that identifies a group of dependencies.
//
// Groups are used to group multiple dependencies of different types that are
// related in some way.
//
// Groups are declared by declaring a named type that uses imbue.Group as its
// underlying type.
type Group interface {
	group()
}

// FromGroup declares a dependency on a type within a specific group.
//
// It is used as a parameter type within user-defined functions passed to
// WithX(), DecorateX(), and InvokeX() to request of type T that is within the
// group G.
type FromGroup[G Group, T any] struct {
	value T
}

// Group returns the name given to the group.
func (v FromGroup[G, T]) Group() string {
	return typeOf[G]().Name()
}

// Value returns the dependency value.
func (v FromGroup[G, T]) Value() T {
	return v.value
}

// inGroup wraps a value of type T to present it as a FromGroup[G, T].
func inGroup[G Group, T any](v T) FromGroup[G, T] {
	return FromGroup[G, T]{
		value: v,
	}
}

// WithGroupedOption is an option that changes the behavior of a call to
// WithXGrouped().
type WithGroupedOption interface {
	applyWithGroupedOption()
}
