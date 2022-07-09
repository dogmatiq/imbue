package lifecycle_test

import (
	"context"
	"errors"

	. "github.com/dogmatiq/imbue/internal/lifecycle"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Private", func() {
	var (
		policy  Private[int]
		factory Factory[int]
		count   int
	)

	BeforeEach(func() {
		policy = Private[int]{}

		factory = func(ctx context.Context) (*Instance[int], error) {
			count++

			return &Instance[int]{
				Value: count,
				Releaser: func() error {
					count--
					return nil
				},
			}, nil
		}

		count = 0
	})

	AfterEach(func() {
		err := policy.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("func Acquire()", func() {
		It("returns a new instance on each call", func() {
			inst, err := policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
			defer inst.Release()

			Expect(inst.Value).To(Equal(1))

			inst, err = policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
			defer inst.Release()

			Expect(inst.Value).To(Equal(2))
		})

		It("uses the factory's releaser on each instance", func() {
			inst, err := policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())

			err = inst.Release()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(0))
		})

		It("returns an error if the factory returns an error", func() {
			factory = func(ctx context.Context) (*Instance[int], error) {
				return nil, errors.New("<error>")
			}

			_, err := policy.Acquire(context.Background(), factory)
			Expect(err).To(MatchError("<error>"))
		})
	})

	Describe("func Close()", func() {
		It("does nothing", func() {
			err := policy.Close()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
