package lifecycle

import (
	"context"
	"sync"
	"sync/atomic"
)

// Singleton is a lifecycle policy that reuses a single value of type T that is
// released when the container is closed.
type Singleton[T any] struct {
	resolved uint32
	m        sync.Mutex
	value    T
	release  Releaser
}

var _ Policy[int] = (*Singleton[int])(nil)

func (p *Singleton[T]) Acquire(
	ctx context.Context,
	new Factory[T],
) (T, Releaser, error) {
	if atomic.LoadUint32(&p.resolved) == 0 {
		p.m.Lock()
		defer p.m.Unlock()

		if p.resolved == 0 {
			value, release, err := new(ctx)
			if err != nil {
				return value, nil, err
			}

			p.value = value
			p.release = release
			atomic.StoreUint32(&p.resolved, 1)
		}
	}

	return p.value, noopReleaser, nil
}

func (p *Singleton[T]) Close() error {
	if atomic.LoadUint32(&p.resolved) == 0 {
		p.m.Lock()
		defer p.m.Unlock()

		if p.resolved == 0 {
			return nil
		}
	}

	zero(&p.value)

	return p.release()
}
