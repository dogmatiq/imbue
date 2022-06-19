package imbue

// InvokeOption is an option that changes the behavior of a call to InvokeX().
type InvokeOption interface {
	applyInvokeOptionToContainer(*Container) error
	applyInvokeOptionToContext(*Context) error
}
