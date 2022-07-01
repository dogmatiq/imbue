package imbue

import (
	"context"
	"reflect"
)

// Context is the context used during construction of dependencies within a
// container.
type Context struct {
	context.Context

	deferrer *deferrer
	funcType string
	declType reflect.Type
}

// Defer registers a function to be invoked when the container is closed.
func (c *Context) Defer(fn func() error) {
	c.deferrer.Add(
		c.funcType,
		c.declType,
		fn,
	)
}

// newChild returns a new context with c as its parent.
func (c *Context) newChild(
	f string,
	t reflect.Type,
) *Context {
	return &Context{
		Context:  c,
		deferrer: c.deferrer,
		funcType: f,
		declType: t,
	}
}

// rootContext returns a new root context.
func rootContext(ctx context.Context, con *Container) *Context {
	return &Context{
		Context:  ctx,
		deferrer: &con.deferrer,
	}
}
