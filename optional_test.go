package imbue_test

import (
	"context"
	"errors"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Optional", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	It("enables the container to construct values of the declared type", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete>", nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep imbue.Optional[Concrete1],
			) error {
				v, err := dep.Value()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(v).To(Equal(Concrete1("<concrete>")))
				return nil
			},
		)
	})

	It("treats the dependency as unavailable if no constructor is declared", func() {
		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep imbue.Optional[Concrete1],
			) error {
				_, err := dep.Value()
				Expect(err).To(MatchError("no constructor is declared for imbue_test.Concrete1"))
				return nil
			},
		)
	})

	It("treats the dependency as unavailable if its constructor returns an error", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "", errors.New("<error>")
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep imbue.Optional[Concrete1],
			) error {
				_, err := dep.Value()
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp(
					`constructor for imbue_test\.Concrete1 \(optional_test\.go:\d+\) failed: <error>`,
				))
				return nil
			},
		)
	})

	It("panics when a cyclic dependency is introduced within a single declaration", func() {
		Expect(func() {
			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					dep imbue.Optional[Concrete1],
				) (Concrete1, error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`constructor for imbue_test\.Concrete1 \(optional_test\.go:\d+\) depends on itself`,
				),
			),
		)
	})

	It("panics when a cyclic dependency is introduced across multiple declarations", func() {
		imbue.With1(
			container,
			func(
				ctx *imbue.Context,
				dep Concrete3,
			) (Concrete1, error) {
				panic("unexpected call")
			},
		)

		imbue.With1(
			container,
			func(
				ctx *imbue.Context,
				dep imbue.Optional[Concrete1],
			) (Concrete2, error) {
				panic("unexpected call")
			},
		)

		Expect(func() {
			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete2,
				) (Concrete3, error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`(?m)constructor for imbue_test\.Concrete3 introduces a cyclic dependency:` +
						`\n\t-> imbue_test\.Concrete2 \(optional_test\.go:\d+\)` +
						`\n\t-> imbue_test\.Concrete1 \(optional_test\.go:\d+\)` +
						`\n\t-> imbue_test\.Concrete3 \(optional_test\.go:\d+\)`,
				),
			),
		)
	})

	It("panics if a constructor is declared for an optional type", func() {
		Expect(func() {
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (imbue.Optional[Concrete1], error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`declaration of constructor for imbue\.Optional\[imbue\.Concrete1\] \(optional_test\.go:\d+\) is disallowed`,
				),
			),
		)
	})
})
