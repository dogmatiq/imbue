package imbue

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

// Context is the context used during initialization of dependencies within a
// container.
type Context struct {
	context.Context
	parent *Context
	entry  *entry
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

// childContext returns a new child context with ctx as its parent.
func childContext(
	ctx *Context,
	e *entry,
) *Context {
	return &Context{
		Context: ctx.Context,
		parent:  ctx,
		entry:   e,
	}
}

// checkForCycle panics if ctx contains a cyclic dependency.
func checkForCycle(ctx *Context, e *entry) {
	stack := []string{
		fmt.Sprintf(
			"container has cyclic dependency involving %s:",
			e.Type.String(),
		),
	}

	for ctx != nil && ctx.parent != nil {
		stack = append(
			stack,
			fmt.Sprintf(
				"%s (%s:%d)",
				ctx.entry.Type.String(),
				filepath.Base(ctx.entry.File),
				ctx.entry.Line,
			),
		)

		if ctx.entry == e {
			panic(strings.Join(stack, "\n\t-> "))
		}

		ctx = ctx.parent
	}
}
