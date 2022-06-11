package imbue_test

import (
	"context"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func With0()", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	It("registers a dependency with the container", func() {
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
})
