package imbue

import (
	"context"
)

// Context is the context used during initialization of dependencies within a
// container.
type Context struct {
	context.Context
}

// Defer registers a function to be invoked when the container is closed.
func (c *Context) Defer(func() error) {
	panic("not implemented")
}

// rootContext returns a new root context.
func rootContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
	}
}
