package imbue

import (
	"context"
)

// Context is the context used during construction of dependencies within a
// container.
type Context struct {
	context.Context

	con    *Container
	scope  userFunction
	defers *deferSet
}

// Defer registers a function to be invoked when the container is closed.
func (c *Context) Defer(fn func() error) {
	c.defers.Add(
		deferred{
			fn,
			findLocation(),
			c.scope,
		},
	)
}

// invokeContext returns a new Context for a function invoked by InvokeX().
func invokeContext(
	parent context.Context,
	con *Container,
) *Context {
	return &Context{
		Context: parent,
		con:     con,
	}
}

// childContext returns a new child Context for use within a constructor or
// decorator.
func childContext(
	parent *Context,
	scope userFunction,
	defers *deferSet,
) *Context {
	return &Context{
		Context: parent,
		con:     parent.con,
		scope:   scope,
		defers:  defers,
	}
}
