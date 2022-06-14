package imbue

import (
	"fmt"
	"reflect"
	"sync"
)

// Container is a dependency injection container.
type Container struct {
	m            sync.Mutex
	declarations map[reflect.Type]declaration
	closers      []func() error
}

// closeError is returned when there are one or more errors closing the
// container.
type closeError []error

func (e closeError) Error() string {
	message := fmt.Sprintf(
		"%d error(s) occurred while closing the container:",
		len(e),
	)

	for _, err := range e {
		message += fmt.Sprintf("\n\t%s", err)
	}

	return message
}

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	c.m.Lock()
	defer c.m.Unlock()

	closers := c.closers
	c.closers = nil

	var errors closeError

	for i := len(closers) - 1; i >= 0; i-- {
		if err := closers[i](); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// addDefer registers a function to be called when the container is closed.
func (c *Container) addDefer(fn func() error) {
	c.m.Lock()
	defer c.m.Unlock()

	c.closers = append(c.closers, fn)
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
