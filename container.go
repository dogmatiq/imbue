package imbue

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/xlab/treeprint"
	"golang.org/x/exp/slices"
)

// Container is a dependency injection container.
type Container struct {
	m            sync.Mutex
	declarations map[reflect.Type]declaration
	deferred     []func() error
}

// closeError is returned when there are one or more errors closing the
// container.
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

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	c.m.Lock()
	defer c.m.Unlock()

	deferred := c.deferred
	c.deferred = nil

	var errors closeError

	for i := len(deferred) - 1; i >= 0; i-- {
		if err := deferred[i](); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// addDefer registers a function to be called when the container is closed.
func (c *Container) addDefer(fn func() error) {
	c.m.Lock()
	defer c.m.Unlock()

	c.deferred = append(c.deferred, fn)
}

// New returns a new, empty container.
func New() *Container {
	return &Container{
		declarations: map[reflect.Type]declaration{},
	}
}

// typeOf returns the reflect.Type for T.
func typeOf[T any]() reflect.Type {
	return reflect.TypeOf([0]T{}).Elem()
}

// get returns the declaration for type T.
func get[T any](con *Container) *declarationOf[T] {
	t := typeOf[T]()

	con.m.Lock()
	defer con.m.Unlock()

	if d, ok := con.declarations[t]; ok {
		return d.(*declarationOf[T])
	}

	d := &declarationOf[T]{}
	con.declarations[t] = d

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
		t.AddNode(d.GetType().String())
		return
	}

	sub := t.AddBranch(d.GetType().String())

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

	slices.SortFunc(sorted, func(a, b declaration) bool {
		return a.GetType().String() < b.GetType().String()
	})

	return sorted
}
