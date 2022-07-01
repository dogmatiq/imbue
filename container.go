package imbue

import (
	"reflect"
	"sync"

	"github.com/xlab/treeprint"
	"golang.org/x/exp/slices"
)

// Container is a dependency injection container.
type Container struct {
	m            sync.Mutex
	declarations map[reflect.Type]declaration
	deferrer     deferrer
}

// New returns a new, empty container.
func New() *Container {
	return &Container{
		declarations: map[reflect.Type]declaration{},
	}
}

// Close closes the container, calling any deferred functions registered
// during construction of dependencies.
func (c *Container) Close() error {
	return c.deferrer.Close()
}

// typeOf returns the reflect.Type for T.
func typeOf[T any]() reflect.Type {
	return reflect.TypeOf([0]T{}).Elem()
}

// get returns the declaration for type T.
func get[T any](con *Container) (*declarationOf[T], error) {
	t := typeOf[T]()

	con.m.Lock()

	if d, ok := con.declarations[t]; ok {
		con.m.Unlock()
		return d.(*declarationOf[T]), nil
	}

	d := &declarationOf[T]{}
	con.declarations[t] = d

	con.m.Unlock()

	if err := d.Init(con); err != nil {
		return nil, err
	}

	return d, nil
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
