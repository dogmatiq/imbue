package lifecycle

import (
	"context"
	"sync"
)

// Shared is a lifecycle policy that reuses a single value of type T across a
// set of concurrent InvokeX() calls.
//
// The value is released when there are no concurrent InvokeX() calls that
// depend on T.
type Shared[T any] struct {
	m       sync.Mutex
	refs    int
	value   T
	release Releaser
}

var _ Policy[int] = (*Shared[int])(nil)

func (p *Shared[T]) Acquire(
	ctx context.Context,
	new Factory[T],
) (T, Releaser, error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.refs == 0 {
		value, release, err := new(ctx)
		if err != nil {
			return value, nil, err
		}

		p.value = value
		p.release = release
	}

	p.refs++

	return p.value, p.releaseOne, nil
}

func (p *Shared[T]) releaseOne() error {
	p.m.Lock()
	defer p.m.Unlock()

	p.refs--
	if p.refs > 0 {
		return nil
	}

	defer func() {
		zero(&p.value)
		p.release = nil
	}()

	return p.release()
}

func (p *Shared[T]) Close() error {
	return nil
}
