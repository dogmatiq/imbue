package lifecycle

import (
	"context"
)

// Policy is an interface for a lifecycle policy, which defines the lifecycle of
// a value that is obtained from a container.
type Policy[T any] interface {
	// Acquire obtains an instance of type T.
	Acquire(context.Context, Factory[T]) (*Instance[T], error)

	// Close closes the policy, releasing any resources that it may be holding.
	Close() error
}
