package imbue

import (
	"fmt"
	"path/filepath"
	"reflect"
	"sync"
)

// declaration is an interface that describes how to build a dependency.
type declaration interface {
}

type constructor[T any] func(*Context) (T, error)

// declarationOf describes how to build dependencies of a specific type.
type declarationOf[T any] struct {
	m sync.Mutex

	file string
	line int

	isConstructed bool
	dependencies  map[reflect.Type]struct{}
	construct     constructor[T]
	value         T
}

func (d *declarationOf[T]) Declare(
	file string,
	line int,
	decl func() (constructor[T], error),
) error {
	d.m.Lock()
	defer d.m.Unlock()

	d.file = file
	d.line = line

	c, err := decl()
	if err != nil {
		return err
	}

	d.construct = c
	return nil
}

func (d *declarationOf[T]) Resolve(ctx *Context) (T, error) {
	d.m.Lock()
	defer d.m.Unlock()

	if d.isConstructed {
		return d.value, nil
	}

	if d.construct == nil {
		panic(fmt.Sprintf(
			"no constructor is declared for %s",
			typeOf[T](),
		))
	}

	v, err := d.construct(ctx)
	if err != nil {
		return d.value, fmt.Errorf(
			"constructor for %s (%s:%d) failed: %w",
			typeOf[T](),
			filepath.Base(d.file),
			d.line,
			err,
		)
	}

	d.isConstructed = true
	d.dependencies = nil
	d.construct = nil
	d.value = v

	return v, nil
}

// addDependency declares that T has a dependency on type D and returns an
// error if it introduces a cyclic dependency.
func addDependency[T, D any](
	decl *declarationOf[T],
	dep *declarationOf[D],
) error {
	return nil
}
