package imbue

import "errors"

// InvokeOption is an option that changes the behavior of a call to InvokeX().
type InvokeOption interface {
	applyInvokeOptionToContainer(*Container) error
	applyInvokeOptionToContext(*scopedContext) error
}

// filterInvokeError is called when an InvokeX() function is about to return an
// error.
//
// It may modify the error, return a new error, or panic.
func filterInvokeError(err error) error {
	var u undeclaredConstructorError
	if errors.As(err, &u) {
		panic(u.Error())
	}

	return err
}
