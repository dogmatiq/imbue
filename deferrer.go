package imbue

import (
	"fmt"
	"reflect"
	"sync"
)

// deferrer is a registry of deferred functions that are to be called when a
// container is closed.
type deferrer struct {
	m      sync.Mutex
	defers []deferEntry
}

type deferEntry struct {
	FuncType string
	DeclType reflect.Type
	Location location
	Func     func() error
}

// Add registers a function to be called when the deferrer is closed.
func (d *deferrer) Add(
	f string,
	t reflect.Type,
	fn func() error,
) {
	loc := findLocation()

	d.m.Lock()
	defer d.m.Unlock()

	d.defers = append(d.defers, deferEntry{
		FuncType: f,
		DeclType: t,
		Func:     fn,
		Location: loc,
	})
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

		// All the deferred functions using an actual defer statement, thus
		// guaranteeing that the functions are invoked in reverse order _and_
		// that they are always invoked, even if one of them panics.
		defer func() {
			if err := e.Func(); err != nil {
				errors = append(
					errors,
					fmt.Errorf(
						"deferred by %s %s at %s: %w",
						e.DeclType,
						e.FuncType,
						e.Location,
						err,
					),
				)
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
