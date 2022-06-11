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

	It("can obtain dependencies from the container", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete>", nil
			},
		)

		imbue.With1(
			container,
			func(ctx *imbue.Context, dep Concrete1) (Concrete2, error) {
				return Concrete2(dep), nil
			},
		)

		imbue.InvokeWith1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete2,
			) error {
				Expect(dep).To(Equal(Concrete2("<concrete>")))
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
