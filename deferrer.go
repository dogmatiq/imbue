package imbue

import (
	"fmt"
	"sync"
)

// deferrer is a registry of deferred functions that are to be called when a
// container is closed.
type deferrer struct {
	m     sync.Mutex
	funcs []func() error
}

// Add registers a function to be called when the deferrer is closed.
func (d *deferrer) Add(fn func() error) {
	d.m.Lock()
	defer d.m.Unlock()

	d.funcs = append(d.funcs, fn)
}

// Close invokes the registered functions in reverse order.
func (d *deferrer) Close() error {
	d.m.Lock()
	defer d.m.Unlock()

	if errors := d.call(); len(errors) != 0 {
		return deferError(errors)
	}

	return nil
}

func (d *deferrer) call() (errors []error) {
	funcs := d.funcs
	d.funcs = nil

	for _, fn := range funcs {
		fn := fn // capture loop variable

		// All the deferred functions using an actual defer statement, thus
		// guaranteeing that the functions are invoked in reverse order _and_
		// that they are always invoked, even if one of them panics.
		defer func() {
			if err := fn(); err != nil {
				errors = append(errors, err)
			}
		}()
	}

	return errors
}

// deferError is returned when there are one or more errors returned by deferred
// functions.
type deferError []error

func (e deferError) Error() string {
	message := fmt.Sprintf(
		"%d error(s) occurred while calling deferred functions:",
		len(e),
	)

	for i, err := range e {
		message += fmt.Sprintf("\n\t%d) %s", i+1, err)
	}

	return message
}
