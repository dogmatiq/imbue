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

	It("registers a constructor with the container", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete>", nil
			},
		)

		imbue.InvokeWith1(
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

	It("can obtain a single dependency from the container", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete-1>", nil
			},
		)

		imbue.With1(
			container,
			func(ctx *imbue.Context, dep Concrete1) (Concrete2, error) {
				Expect(dep).To(Equal(Concrete1("<concrete-1>")))
				return "<concrete-2>", nil
			},
		)

		imbue.InvokeWith1(
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

	It("can obtain multiple dependencies from the container", func() {
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

		imbue.InvokeWith1(
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

		imbue.InvokeWith1(
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

		imbue.InvokeWith1(
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
})
