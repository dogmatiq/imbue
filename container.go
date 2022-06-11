package imbue

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"
)

// Container is a dependency injection container.
type Container struct {
	m     sync.RWMutex
	types map[reflect.Type]*entry
}

// entry is an entry within a container for a specific type.
type entry struct {
	// IsConstructed indicates whether the dependency has been constructed.
	// If it is true, then Value contains the constructed value.
	IsConstructed bool

	// Value is the constructed value, which may be nil.
	// It is only valid if IsConstructed is true.
	Value any

	// New constructs the value.
	New func(*Context, *Container) (any, error)

	// File is the file name of the source code that registered the dependency.
	File string

	// Line is the line number of the source code that registered the
	// dependency.
	Line int
}

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	return nil
}

// New returns a new, empty container.
func New() *Container {
	return &Container{}
}

func (c *Container) register(
	typ reflect.Type,
	new func(*Context, *Container) (any, error),
) {
	_, file, line, _ := runtime.Caller(3)

	c.m.Lock()
	defer c.m.Unlock()

	if c.types == nil {
		c.types = map[reflect.Type]*entry{}
	}

	c.types[typ] = &entry{
		New:  new,
		File: file,
		Line: line,
	}
}

func (c *Container) get(typ reflect.Type) (any, error) {
	c.m.RLock()
	defer c.m.RUnlock()

	e, ok := c.types[typ]
	if !ok {
		panic(fmt.Sprintf(
			"the container has no constructor registered for dependencies of type '%s'",
			typ,
		))
	}

	if e.IsConstructed {
		return e.Value, nil
	}

	v, err := e.New(nil, c)
	if err != nil {
		return nil, fmt.Errorf(
			"unable to construct dependency of type '%s' (%s:%d): %w",
			typ,
			filepath.Base(e.File),
			e.Line,
			err,
		)
	}

	e.IsConstructed = true
	e.Value = v
	e.New = nil

	return e.Value, nil
}

// register is a helper function for registering a constructor with a type.
func register[T any](
	c *Container,
	fn func(*Context, *Container) (T, error),
) {
	c.register(
		reflect.TypeOf(fn).Out(0),
		func(ctx *Context, con *Container) (any, error) {
			return fn(ctx, con)
		},
	)
}

// get is a helper function for getting typed values out of a container.
func get[T any](c *Container) (result T, _ error) {
	v, err := c.get(reflect.TypeOf(result))
	if err != nil {
		return result, err
	}

	return v.(T), err
}
