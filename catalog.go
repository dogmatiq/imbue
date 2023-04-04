package imbue

import "sync"

// Catalog is a reusable set of constructors and decorators that can be added to
// different containers.
//
// It implements the ContainerAware interface, so it can be passed to the
// WithX(), WithXNamed(), WithXGrouped() and DecorateX() functions instead of a
// container.
type Catalog struct {
	m     sync.RWMutex
	funcs []func(*Container)
}

// CatalogOption is an option that changes the behavior of a Catalog.
type CatalogOption interface {
	applyCatalogOption(*Catalog)
}

// NewCatalog creturns a new catalog.
func NewCatalog(options ...CatalogOption) *Catalog {
	c := &Catalog{}

	for _, opt := range options {
		opt.applyCatalogOption(c)
	}

	return c
}

func (c *Catalog) withContainer(fn func(*Container)) {
	c.m.Lock()
	defer c.m.Unlock()
	c.funcs = append(c.funcs, fn)
}

// WithCatalog is a ContainerOption that adds the declarations in the catalog to
// the container.
func WithCatalog(cat *Catalog) ContainerOption {
	return option{
		forContainer: func(con *Container) {
			cat.m.Lock()
			defer cat.m.Unlock()
			for _, fn := range cat.funcs {
				fn(con)
			}
		},
	}
}
