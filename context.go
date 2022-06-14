package imbue

import (
	"context"
)

// Context is the context used during initialization of dependencies within a
// container.
type Context struct {
	context.Context
	con *Container
}

// Defer registers a function to be invoked when the container is closed.
func (c *Context) Defer(fn func() error) {
	c.con.addDefer(fn)
}

// rootContext returns a new root context.
func rootContext(ctx context.Context, con *Container) *Context {
	return &Context{
		Context: ctx,
		con:     con,
	}
}
