package imbue

import (
	"fmt"
	"sync"
)

// deferrer is a registry of deferred functions that are to be called when a
// container is closed.
type deferrer struct {
	m      sync.Mutex
	defers []deferred
}

// deferred is a wrapper around a function that is deferred during construction
// or decoration of a dependency.
//
// It implements the userFunction interface.
type deferred struct {
	// impl is the deferred function.
	impl func() error

	// loc is the location of the code that deferred the function.
	loc location

	// scope is the constructor or decorator that deferred the function.
	scope userFunction
}

// Call invokes the deferred function.
func (d deferred) Call() error {
	if err := d.impl(); err != nil {
		return fmt.Errorf(
			"%s failed: %w",
			d,
			err,
		)
	}

	return nil
}

// Location returns the location of the code that deferred the function.
//
// This is typically the location of the call to the Context.Defer() method, not
// the deferred function's definition.
func (d deferred) Location() location {
	return d.loc
}

// String returns a description of the deferred function for use in error
// messages.
func (d deferred) String() string {
	return fmt.Sprintf(
		"function deferred at %s by %s",
		d.loc,
		d.scope,
	)
}

// Add registers a function to be called when the deferrer is closed.
func (d *deferrer) Add(df deferred) {
	d.m.Lock()
	defer d.m.Unlock()

	d.defers = append(d.defers, df)
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
