package imbue

// WithOption is an option that changes the behavior of a call to WithX().
type WithOption interface {
	applyWithOption()
}
