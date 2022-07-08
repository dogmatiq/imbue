package lifecycle

import (
	"context"
)

// Private is a lifecycle policy that creates a value of type T for each
// dependent.
type Private[T any] struct{}

var _ Policy[int] = (*Private[int])(nil)

func (p *Private[T]) Acquire(
	ctx context.Context,
	new Factory[T],
) (T, Releaser, error) {
	return new(ctx)
}

func (p *Private[T]) Close() error {
	return nil
}
