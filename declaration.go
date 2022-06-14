package imbue

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"sync"
)

// declaration is an interface that describes how to build a dependency.
type declaration interface {
	AddDependency(t declaration) error

	getType() reflect.Type
	location() (string, int)
	dependsOn(t declaration, cycle []declaration) ([]declaration, bool)
}

// constructor is a function that constructs values of type T.
type constructor[T any] func(*Context) (T, error)

// declarationOf describes how to build dependencies of a specific type.
type declarationOf[T any] struct {
	m sync.Mutex

	file string
	line int

	isConstructed bool
	dependencies  map[reflect.Type]declaration
	construct     constructor[T]
	value         T
}

func (d *declarationOf[T]) Declare(
	file string,
	line int,
	decl func() (constructor[T], error),
) error {
	d.m.Lock()
	d.file = file
	d.line = line
	d.m.Unlock()

	c, err := decl()
	if err != nil {
		return err
	}

	d.m.Lock()
	d.construct = c
	d.m.Unlock()

	return nil
}

func (d *declarationOf[T]) AddDependency(t declaration) error {
	if t == d {
		file, line := d.location()

		return fmt.Errorf(
			"constructor for %s (%s:%d) depends on itself",
			d.getType(),
			filepath.Base(file),
			line,
		)
	}

	if cycle, ok := t.dependsOn(d, nil); ok {
		message := fmt.Sprintf(
			"constructor for %s introduces a cyclic dependency:",
			d.getType(),
		)

		_ = cycle

		for i := len(cycle) - 1; i >= 0; i-- {
			dep := cycle[i]
			file, line := dep.location()

			message += fmt.Sprintf(
				"\n\t-> %s (%s:%d)",
				dep.getType(),
				filepath.Base(file),
				line,
			)
		}

		return errors.New(message)
	}

	d.m.Lock()
	defer d.m.Unlock()

	if d.dependencies == nil {
		d.dependencies = map[reflect.Type]declaration{}
	}

	d.dependencies[t.getType()] = t

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

func (d *declarationOf[T]) getType() reflect.Type {
	return typeOf[T]()
}

func (d *declarationOf[T]) location() (string, int) {
	d.m.Lock()
	defer d.m.Unlock()

	return d.file, d.line
}

func (d *declarationOf[T]) dependsOn(t declaration, cycle []declaration) ([]declaration, bool) {
	if t.getType() == d.getType() {
		return append(cycle, d), true
	}

	d.m.Lock()
	defer d.m.Unlock()

	for _, dep := range d.dependencies {
		if cycle, ok := dep.dependsOn(t, cycle); ok {
			return append(cycle, d), true
		}
	}

	return nil, false
}
