package imbue

import (
	"fmt"
)

// deferSet is a set of deferred functions.
type deferSet struct {
	defers []deferred
}

// Add adds a deferred function to the set.
func (s *deferSet) Add(d deferred) {
	s.defers = append(s.defers, d)
}

// Call invokes the deferred functions in reverse order.
func (s *deferSet) Call() (errors []error) {
	defers := s.defers
	s.defers = nil

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

// TransferOwnership transfers ownership of this set's deferred function to
// target.
func (s *deferSet) TransferOwnership(target *deferSet) {
	target.defers = append(target.defers, s.defers...)
	s.defers = nil
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
