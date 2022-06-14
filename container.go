package imbue

import (
	"reflect"
	"sync"
)

// Container is a dependency injection container.
type Container struct {
	m            sync.Mutex
	declarations map[reflect.Type]declaration
}

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	return nil
}

// New returns a new, empty container.
func New() *Container {
	return &Container{
		declarations: map[reflect.Type]declaration{},
	}
}

// typeOf returns the reflect.Type for T.
func typeOf[T any]() reflect.Type {
	return reflect.TypeOf([0]T{}).Elem()
}

// get returns the declaration for type T.
func get[T any](con *Container) *declarationOf[T] {
	t := typeOf[T]()

	con.m.Lock()
	defer con.m.Unlock()

	if d, ok := con.declarations[t]; ok {
		return d.(*declarationOf[T])
	}

	d := &declarationOf[T]{}
	con.declarations[t] = d

	return d
}
