package imbue

import (
	"fmt"
	"sync"
)

// deferrer is a registry of deferred functions that are to be called when a
// container is closed.
type deferrer struct {
	m      sync.Mutex
	defers []deferEntry
}

type deferEntry struct {
	scope userFunction
	xfn   func() error
}

func (e deferEntry) Call() error {
	if err := e.xfn(); err != nil {
		return fmt.Errorf(
			"deferred by %s: %w",
			e.scope,
			err,
		)
	}

	return nil
}

// Add registers a function to be called when the deferrer is closed.
func (d *deferrer) Add(e deferEntry) {
	d.m.Lock()
	defer d.m.Unlock()

	d.defers = append(d.defers, e)
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
	defers := d.defers
	d.defers = nil

	for _, e := range defers {
		e := e // capture loop variable

		// Call the deferred functions using an actual defer statement, thus
		// guaranteeing that the functions are invoked in reverse order _and_
		// that they are always invoked, even if one of them panics.
		defer func() {
			if err := e.Call(); err != nil {
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
		"%d error(s) occurred in deferred functions:",
		len(e),
	)

	for i, err := range e {
		message += fmt.Sprintf("\n\t%d) %s", i+1, err)
	}

	return message
}
