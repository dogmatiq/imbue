package lifecycle

import (
	"context"
	"sync"
	"sync/atomic"
)

// Singleton is a lifecycle policy that reuses a single value of type T that is
// released when the container is closed.
type Singleton[T any] struct {
	exists uint32 // atomic bool
	m      sync.Mutex
	inst   *Instance[T]
}

var _ Policy[int] = (*Singleton[int])(nil)

func (p *Singleton[T]) Acquire(
	ctx context.Context,
	new Factory[T],
) (*Instance[T], error) {
	if atomic.LoadUint32(&p.exists) == 0 {
		p.m.Lock()
		defer p.m.Unlock()

		if p.exists == 0 {
			inst, err := new(ctx)
			if err != nil {
				return nil, err
			}

			p.inst = inst
			atomic.StoreUint32(&p.exists, 1)
		}
	}

	return &Instance[T]{
		Value: p.inst.Value,
	}, nil
}

func (p *Singleton[T]) Close() error {
	p.m.Lock()
	defer p.m.Unlock()

	if p.exists == 0 {
		return nil
	}

	inst := p.inst

	p.exists = 0
	p.inst = nil

	return inst.Release()
}
