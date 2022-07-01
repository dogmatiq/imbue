package imbue_test

import (
	"context"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func WithX()", func() {
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
				dep Concrete1,
			) error {
				Expect(dep).To(Equal(Concrete1("<concrete>")))
				return nil
			},
		)
	})

	It("can request a single dependency via the constructor's input parameters", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete-1>", nil
			},
		)

		imbue.With1(
			container,
			func(
				ctx *imbue.Context,
				dep Concrete1,
			) (Concrete2, error) {
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
				return "<concrete-2>", nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete2,
			) error {
				Expect(dep).To(Equal(Concrete2("<concrete-2>")))
				return nil
			},
		)
	})

	It("can request multiple dependencies via the constructor's input parameters", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete-1>", nil
			},
		)

		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete2, error) {
				return "<concrete-2>", nil
			},
		)

		imbue.With2(
			container,
			func(
				ctx *imbue.Context,
				dep1 Concrete1,
				dep2 Concrete2,
			) (Concrete3, error) {
				Expect(dep1).To(Equal(Concrete1("<concrete-1>")))
				Expect(dep2).To(Equal(Concrete2("<concrete-2>")))
				return Concrete3("<concrete-3>"), nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete3,
			) error {
				Expect(dep).To(Equal(Concrete3("<concrete-3>")))
				return nil
			},
		)
	})

	It("can request dependencies that have dependencies of their own", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete-1>", nil
			},
		)

		imbue.With1(
			container,
			func(
				ctx *imbue.Context,
				dep Concrete1,
			) (Concrete2, error) {
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
				return "<concrete-2>", nil
			},
		)

		imbue.With1(
			container,
			func(
				ctx *imbue.Context,
				dep Concrete2,
			) (Concrete3, error) {
				Expect(dep).To(Equal(Concrete2("<concrete-2>")))
				return "<concrete-3>", nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete3,
			) error {
				Expect(dep).To(Equal(Concrete3("<concrete-3>")))
				return nil
			},
		)
	})

	It("only invokes the constructor once even if the value is requested multiple times", func() {
		called := false
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				Expect(called).To(BeFalse(), "constructor called multiple times")
				called = true
				return "<concrete>", nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete1,
			) error {
				Expect(dep).To(Equal(Concrete1("<concrete>")))
				return nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete1,
			) error {
				Expect(dep).To(Equal(Concrete1("<concrete>")))
				return nil
			},
		)
	})

	It("panics if the type has already been declared", func() {
		imbue.With0(
			container,
			func(
				ctx *imbue.Context,
			) (Concrete1, error) {
				panic("unexpected call")
			},
		)

		Expect(func() {
			imbue.With0(
				container,
				func(
					ctx *imbue.Context,
				) (Concrete1, error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`imbue_test\.Concrete1 constructor \(with_test\.go:\d+\) collides with existing constructor declared at with_test\.go:\d+`,
				),
			),
		)
	})

	It("panics when a cyclic dependency is introduced within a single declaration", func() {
		Expect(func() {
			imbue.With1(
				container,
				func(
					ctx *imbue.Context,
					dep Concrete1,
				) (Concrete1, error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`imbue_test\.Concrete1 constructor \(with_test\.go:\d+\) depends on itself`,
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
				dep Concrete1,
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
					`(?m)imbue_test\.Concrete3 constructor introduces a cyclic dependency:` +
						`\n\t-> imbue_test\.Concrete2 \(with_test\.go:\d+\)` +
						`\n\t-> imbue_test\.Concrete1 \(with_test\.go:\d+\)` +
						`\n\t-> imbue_test\.Concrete3 \(with_test\.go:\d+\)`,
				),
			),
		)
	})
})
