package imbue

import (
	"context"
)

// Context is the context used during construction of dependencies within a
// container.
type Context struct {
	context.Context

	deferrer *deferrer
	scope    userFunction
}

// Defer registers a function to be invoked when the container is closed.
func (c *Context) Defer(fn func() error) {
	c.deferrer.Add(
		deferred{
			fn,
			findLocation(),
			c.scope,
		},
	)
}

// rootContext returns a new root context.
func rootContext(ctx context.Context, con *Container) *Context {
	return &Context{
		Context:  ctx,
		deferrer: &con.deferrer,
	}
}
