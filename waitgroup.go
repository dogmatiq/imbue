package imbue

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// A WaitGroup is a collection of goroutines working on subtasks that are part
// of the same overall task, each of which has dependencies provided by a
// container.
type WaitGroup struct {
	con   *Container
	ctx   context.Context
	group *errgroup.Group
}

// Container returns the container that the group is bound to.
func (g *WaitGroup) Container() *Container {
	return g.con
}

// Wait blocks until all function calls from the Go method have returned,
// then returns the first non-nil error (if any) from them.
func (g *WaitGroup) Wait() error {
	return g.group.Wait()
}

// func Go1[T any](
// 	g *WaitGroup,
// 	fn func(
// 		ctx context.Context,
// 		dep T,
// 	) error,
// ) {
// 	g.group.Go(func() error {
// 		return Invoke1(
// 			g.ctx,
// 			g.con,
// 			fn,
// 		)
// 	})
// }
