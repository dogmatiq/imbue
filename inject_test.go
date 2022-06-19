package imbue_test

import (
	"context"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func InjectX()", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	It("can request a single dependency via the initializer's input parameters", func() {
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

		called := false
		imbue.Inject1(
			container,
			func(
				ctx *imbue.Context,
				v Concrete2,
				dep Concrete1,
			) error {
				called = true
				Expect(v).To(Equal(Concrete2("<concrete-2>")))
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
				return nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete2,
			) error {
				Expect(called).To(BeTrue())
				Expect(dep).To(Equal(Concrete2("<concrete-2>")))
				return nil
			},
		)

	})

	It("can request multiple dependencies via the initializer's input parameters", func() {
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

		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete3, error) {
				return "<concrete-3>", nil
			},
		)

		called := false
		imbue.Inject2(
			container,
			func(
				ctx *imbue.Context,
				v Concrete3,
				dep1 Concrete1,
				dep2 Concrete2,
			) error {
				called = true
				Expect(v).To(Equal(Concrete3("<concrete-3>")))
				Expect(dep1).To(Equal(Concrete1("<concrete-1>")))
				Expect(dep2).To(Equal(Concrete2("<concrete-2>")))
				return nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete3,
			) error {
				Expect(called).To(BeTrue())
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

		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete3, error) {
				return "<concrete-3>", nil
			},
		)

		called := false
		imbue.Inject1(
			container,
			func(
				ctx *imbue.Context,
				v Concrete3,
				dep Concrete2,
			) error {
				called = true
				Expect(v).To(Equal(Concrete3("<concrete-3>")))
				Expect(dep).To(Equal(Concrete2("<concrete-2>")))
				return nil
			},
		)

		imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete3,
			) error {
				Expect(called).To(BeTrue())
				Expect(dep).To(Equal(Concrete3("<concrete-3>")))
				return nil
			},
		)
	})

	It("only invokes the initializer once even if the value is requested multiple times", func() {
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

		called := false
		imbue.Inject1(
			container,
			func(
				ctx *imbue.Context,
				v Concrete1,
				dep Concrete2,
			) error {
				Expect(called).To(BeFalse(), "constructor called multiple times")
				called = true
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
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
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
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
				return nil
			},
		)
	})

	It("panics when a cyclic dependency is introduced within a single declaration", func() {
		Expect(func() {
			imbue.Inject1(
				container,
				func(
					ctx *imbue.Context,
					v Concrete1,
					dep Concrete1,
				) error {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchError(
					MatchRegexp(
						`initializer for imbue_test\.Concrete1 \(inject_test\.go:\d+\) depends on itself`,
					),
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
			imbue.Inject1(
				container,
				func(
					ctx *imbue.Context,
					v Concrete3,
					dep Concrete2,
				) error {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchError(
					MatchRegexp(
						`(?m)initializer for imbue_test\.Concrete3 introduces a cyclic dependency:` +
							`\n\t-> imbue_test\.Concrete2 \(inject_test\.go:\d+\)` +
							`\n\t-> imbue_test\.Concrete1 \(inject_test\.go:\d+\)` +
							`\n\t-> imbue_test\.Concrete3 \(inject_test\.go:\d+\)`,
					),
				),
			),
		)
	})
})