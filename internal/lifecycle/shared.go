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
	m    sync.Mutex
	refs int
	inst *Instance[T]
}

var _ Policy[int] = (*Shared[int])(nil)

func (p *Shared[T]) Acquire(
	ctx context.Context,
	new Factory[T],
) (*Instance[T], error) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.refs == 0 {
		inst, err := new(ctx)
		if err != nil {
			return nil, err
		}

		p.inst = inst
	}

	p.refs++

	return &Instance[T]{
		p.inst.Value,
		p.releaseOne,
	}, nil
}

func (p *Shared[T]) releaseOne() error {
	p.m.Lock()
	defer p.m.Unlock()

	p.refs--
	if p.refs > 0 {
		return nil
	}

	defer func() {
		p.inst = nil
	}()

	return p.inst.Release()
}

func (p *Shared[T]) Close() error {
	return nil
}
