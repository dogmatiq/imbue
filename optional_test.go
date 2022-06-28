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
})
