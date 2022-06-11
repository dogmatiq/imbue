package imbue_test

import (
	"context"
	"errors"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func InvokeX()", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	It("can obtain a single dependency from the container", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete-1>", nil
			},
		)

		err := imbue.InvokeWith1(
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
		Expect(err).ShouldNot(HaveOccurred())
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

		err := imbue.InvokeWith2(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep1 Concrete1,
				dep2 Concrete2,
			) error {
				Expect(dep1).To(Equal(Concrete1("<concrete-1>")))
				Expect(dep2).To(Equal(Concrete2("<concrete-2>")))
				return nil
			},
		)
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("returns the error returned by the invoked function", func() {
		imbue.With0(
			container,
			func(ctx *imbue.Context) (Concrete1, error) {
				return "<concrete>", nil
			},
		)

		err := imbue.InvokeWith1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				dep Concrete1,
			) error {
				return errors.New("<error>")
			},
		)
		Expect(err).To(MatchError("<error>"))
	})
})
