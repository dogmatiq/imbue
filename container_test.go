package imbue_test

import (
	"strings"

	"github.com/dogmatiq/imbue"
	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("type Container", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	Describe("func String()", func() {
		It("renders a single root", func() {
			imbue.With0(
				container,
				func(
					*imbue.Context,
				) (Concrete1, error) {
					panic("not implemented")
				},
			)

			expectMultilineString(
				container,
				"<container>",
				"└── imbue_test.Concrete1",
			)
		})

		It("renders a single root with a dependency", func() {
			imbue.With1(
				container,
				func(
					*imbue.Context,
					Concrete2,
				) (Concrete1, error) {
					panic("not implemented")
				},
			)

			expectMultilineString(
				container,
				"<container>",
				"└── imbue_test.Concrete1",
				"    └── imbue_test.Concrete2",
			)
		})

		It("renders a single root with multiple dependencies", func() {
			imbue.With2(
				container,
				func(
					*imbue.Context,
					Concrete2,
					Concrete3,
				) (Concrete1, error) {
					panic("not implemented")
				},
			)

			expectMultilineString(
				container,
				"<container>",
				"└── imbue_test.Concrete1",
				"    ├── imbue_test.Concrete2",
				"    └── imbue_test.Concrete3",
			)
		})

		It("renders a single root with a chain of dependencies", func() {
			imbue.With1(
				container,
				func(
					*imbue.Context,
					Concrete2,
				) (Concrete1, error) {
					panic("not implemented")
				},
			)

			imbue.With1(
				container,
				func(
					*imbue.Context,
					Concrete3,
				) (Concrete2, error) {
					panic("not implemented")
				},
			)

			expectMultilineString(
				container,
				"<container>",
				"└── imbue_test.Concrete1",
				"    └── imbue_test.Concrete2",
				"        └── imbue_test.Concrete3",
			)
		})
	})
})

func expectMultilineString(
	container *imbue.Container,
	expected ...string,
) {
	expected = append(expected, "")
	actual := strings.Split(container.String(), "\n")

	if diff := cmp.Diff(expected, actual); diff != "" {
		Fail(diff)
	}
}
