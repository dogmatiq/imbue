package imbue

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/xlab/treeprint"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

// ContainerAware is an interface for types that can operate on a container.
type ContainerAware interface {
	withContainer(func(*Container))
}

// Container is a dependency injection container.
type Container struct {
	m            sync.Mutex
	declarations map[reflect.Type]declaration
	defers       deferSet
}

// ContainerOption is an option that changes the behavior of a container or how
// it is constructed.
type ContainerOption interface {
	applyContainerOption(*Container)
}

// New returns a new, empty container.
func New(options ...ContainerOption) *Container {
	con := &Container{
		declarations: map[reflect.Type]declaration{},
	}

	for _, opt := range options {
		opt.applyContainerOption(con)
	}

	return con
}

// WaitGroup returns a new WaitGroup that is bound to this container.
func (c *Container) WaitGroup(ctx context.Context) *WaitGroup {
	g, ctx := errgroup.WithContext(ctx)

	return &WaitGroup{
		con:   c,
		ctx:   ctx,
		group: g,
	}
}

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	c.m.Lock()
	defer c.m.Unlock()

	if errors := c.defers.Call(); len(errors) != 0 {
		return closeError(errors)
	}

	return nil
}

func (c *Container) withContainer(fn func(*Container)) {
	fn(c)
}

// typeOf returns the reflect.Type for T.
func typeOf[T any]() reflect.Type {
	return reflect.TypeOf([0]T{}).Elem()
}

// get returns the declaration for type T.
func get[T any](con *Container) *declarationOf[T] {
	t := typeOf[T]()

	con.m.Lock()

	if d, ok := con.declarations[t]; ok {
		con.m.Unlock()
		return d.(*declarationOf[T])
	}

	d := &declarationOf[T]{
		defers: &con.defers,
	}
	con.declarations[t] = d

	con.m.Unlock()

	d.Init(con)

	return d
}

// String returns a string representation of the dependency tree.
func (c *Container) String() string {
	c.m.Lock()
	declarations := sortDeclarations(c.declarations)
	c.m.Unlock()

	tree := treeprint.New()
	tree.SetValue("<container>")

	for _, d := range declarations {
		if !d.IsDependency() {
			buildTree(tree, d)
		}
	}

	return tree.String()
}

// buildTree builds the tree of dependencies for the given declaration.
func buildTree(t treeprint.Tree, d declaration) {
	dependencies := d.Dependencies()

	if len(dependencies) == 0 {
		t.AddNode(d.Type().String())
		return
	}

	sub := t.AddBranch(d.Type().String())

	for _, dep := range dependencies {
		buildTree(sub, dep)
	}
}

// sortDeclarations returns the given declarations sorted by type.
func sortDeclarations(declarations map[reflect.Type]declaration) []declaration {
	sorted := make([]declaration, 0, len(declarations))

	for _, d := range declarations {
		sorted = append(sorted, d)
	}

	slices.SortFunc(sorted, func(a, b declaration) int {
		return strings.Compare(a.Type().String(), b.Type().String())
	})

	return sorted
}

// closeError is returned when there are one or more errors closing a container.
type closeError []error

func (e closeError) Error() string {
	message := fmt.Sprintf(
		"%d error(s) occurred while closing the container:",
		len(e),
	)

	for i, err := range e {
		message += fmt.Sprintf("\n\t%d) %s", i+1, err)
	}

	return message
}
