package imbue

// WithNamedOption is an option that changes the behavior of a call to
// WithXNamed().
type WithNamedOption interface {
	applyWithNamedOption()
}
