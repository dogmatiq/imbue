package imbue

// InjectOption is an option that changes the behavior of a call to InjectX().
type InjectOption interface {
	applyInjectOption()
}
