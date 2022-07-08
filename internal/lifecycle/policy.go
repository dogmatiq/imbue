package lifecycle

import (
	"context"
)

// Policy is an interface for a lifecycle policy, which defines the lifecycle of
// a value that is obtained from a container.
type Policy[T any] interface {
	// Acquire obtains a reference to the value.
	Acquire(context.Context, Factory[T]) (T, Releaser, error)

	// Close closes the policy, releasing any resources that it may be holding.
	Close() error
}

// Factory is a function that creates new values of type T.
type Factory[T any] func(ctx context.Context) (T, Releaser, error)

// Releaser is a function that releases a reference to a value.
type Releaser func() error

// noopReleaser is a Releaser function that does nothing.
func noopReleaser() error {
	return nil
}

// zero sets *v to its zero-value.
func zero[T any](v *T) {
	var zero T
	*v = zero
}
