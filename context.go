package imbue

import (
	"context"
)

// Context is an extended version of the standard context.Context interface that
// is used when constructing and decorating dependencies.
type Context interface {
	context.Context

	// Defer registers a function to be invoked when the container is closed.
	Defer(fn func() error)
}

// scopedContext is the context used during construction of dependencies within
// a container.
type scopedContext struct {
	context.Context

	scope  userFunction
	defers *deferSet
}

// Defer registers a function to be invoked when the container is closed.
func (c *scopedContext) Defer(fn func() error) {
	c.defers.Add(
		deferred{
			fn,
			findLocation(),
			c.scope,
		},
	)
}
