package lifecycle

import "context"

// Factory is a function that creates new instances of type T.
type Factory[T any] func(ctx context.Context) (*Instance[T], error)

// Decorator is a function that mutates an instance of type T.
type Decorator[T any] func(context.Context, *Instance[T]) error

// Instance is a single instance of a value of type T.
type Instance[T any] struct {
	Value    T
	Releaser func() error
}

func (i Instance[T]) Release() error {
	if i.Releaser == nil {
		return nil
	}

	return i.Releaser()
}
