package imbue

import (
	"context"
	"fmt"
	"reflect"

	"github.com/dogmatiq/imbue/internal/lifecycle"
	locationX "github.com/dogmatiq/imbue/internal/location"
)

// definition is a definition of a specific type within a container.
type definition interface {
	// Type returns the type of the value.
	Type() reflect.Type

	// Location returns the "best" known location of the definition in the
	// user's codebase.
	//
	// This is often (though not always) the location of the WithX() call that
	// declares the constructor for this type.
	Location() locationX.Location
}

type definitionOf[T any] struct {
	location   locationX.Location
	lifecycle  lifecycle.Policy[T]
	factory    lifecycle.Factory[T]
	decorators []lifecycle.Decorator[T]
}

func (d *definitionOf[T]) Type() reflect.Type {
	return typeOf[T]()
}

func (d *definitionOf[T]) Location() locationX.Location {
	return d.location
}

func (d *definitionOf[T]) Acquire(ctx context.Context) (*lifecycle.Instance[T], error) {
	if d.factory == nil {
		return nil, undefinedError{d.Type()}
	}

	return d.lifecycle.Acquire(ctx, d.new)
}

func (d *definitionOf[T]) new(ctx context.Context) (*lifecycle.Instance[T], error) {
	inst, err := d.factory(ctx)
	if err != nil {
		return nil, err
	}

	for _, dec := range d.decorators {
		if err := dec(ctx, inst); err != nil {
			inst.Release()
			return nil, err
		}
	}

	return inst, nil
}

// selfDefining is an interface for types that create their own definitions.
type selfDefining[T any] interface {
	// define initializes the definition of this type.
	//
	// The receiver of this method must be *T, and it will always be nil.
	define(*Container, *definitionOf[T])
}

// def returns the definition of type T within the given container.
func def[T any](con *Container) *definitionOf[T] {
	t := typeOf[T]()

	if def, ok := con.definitions[t]; ok {
		return def.(*definitionOf[T])
	}

	def := &definitionOf[T]{
		location: locationX.Find(),
	}

	con.definitions[t] = def

	if sd, ok := any((*T)(nil)).(selfDefining[T]); ok {
		sd.define(con, def)
	}

	return def
}

// undefinedError is an error returned by definitionOf[T].Acquire() when the
// type has not been defined by a call to one of the WithX() functions.
type undefinedError struct {
	Type reflect.Type
}

func (e undefinedError) Error() string {
	return fmt.Sprintf(
		"no constructor has been defined for %s",
		e.Type,
	)
}
