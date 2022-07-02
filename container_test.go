package imbue_test

import (
	"context"
	"errors"
	"strings"

	"github.com/dogmatiq/imbue"
	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Container", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	Describe("func Close()", func() {
		It("calls deferred functions in reverse order", func() {
			var order []string

			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						order = append(order, "<defer-1>")
						return nil
					})
					return "<concrete-1>", nil
				},
			)

			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					_ Concrete1,
				) (Concrete2, error) {
					ctx.Defer(func() error {
						order = append(order, "<defer-2>")
						return nil
					})
					return "<concrete-2>", nil
				},
			)

			imbue.Invoke1(
				context.Background(),
				container,
				func(
					ctx context.Context,
					_ Concrete2,
				) error {
					return nil
				},
			)

			err := container.Close()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(order).To(Equal([]string{
				"<defer-2>",
				"<defer-1>",
			}))
		})

		It("returns all errors returned by deferred functions", func() {
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						return errors.New("<error-1>")
					})
					return "<concrete-1>", nil
				},
			)

			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					_ Concrete1,
				) (Concrete2, error) {
					ctx.Defer(func() error {
						return errors.New("<error-2>")
					})
					return "<concrete-2>", nil
				},
			)

			imbue.Invoke1(
				context.Background(),
				container,
				func(
					ctx context.Context,
					_ Concrete2,
				) error {
					return nil
				},
			)

			err := container.Close()
			Expect(err).To(
				MatchError(
					MatchRegexp(
						`2 error\(s\) occurred in deferred functions:`+
							`\n\t1\) deferred by imbue_test\.Concrete2 constructor \(container_test\.go:\d+\): <error-2>`+
							`\n\t2\) deferred by imbue_test\.Concrete1 constructor \(container_test\.go:\d+\): <error-1>`,
					),
				),
				err.Error(),
			)
		})

		It("calls all deferred functions even if one of them panics", func() {
			var order []string

			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					ctx.Defer(func() error {
						order = append(order, "<defer-1>")
						return nil
					})
					return "<concrete-1>", nil
				},
			)

			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					_ Concrete1,
				) (Concrete2, error) {
					ctx.Defer(func() error {
						order = append(order, "<defer-2>")
						panic("<panic>")
					})
					return "<concrete-2>", nil
				},
			)

			imbue.Invoke1(
				context.Background(),
				container,
				func(
					ctx context.Context,
					_ Concrete2,
				) error {
					return nil
				},
			)

			Expect(func() {
				container.Close()
			}).To(PanicWith("<panic>"))

			Expect(order).To(Equal([]string{
				"<defer-2>",
				"<defer-1>",
			}))
		})
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
