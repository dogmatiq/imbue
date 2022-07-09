package lifecycle_test

import (
	"context"
	"errors"

	. "github.com/dogmatiq/imbue/internal/lifecycle"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Shared", func() {
	var (
		policy  Shared[int]
		factory Factory[int]
		exists  bool
		count   int
	)

	BeforeEach(func() {
		policy = Shared[int]{}

		factory = func(ctx context.Context) (*Instance[int], error) {
			if exists {
				panic("unexpected call")
			}

			exists = true
			count++

			return &Instance[int]{
				Value: count,
				Releaser: func() error {
					exists = false
					return nil
				},
			}, nil
		}

		exists = false
		count = 0
	})

	AfterEach(func() {
		err := policy.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	When("no value has been acquired", func() {
		Describe("func Acquire()", func() {
			It("returns a new value", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer inst.Release()

				Expect(inst.Value).To(Equal(1))
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

	When("a value has already been acquired", func() {
		var firstInstance *Instance[int]

		BeforeEach(func() {
			var err error
			firstInstance, err = policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			firstInstance.Release()
		})

		Describe("func Acquire()", func() {
			It("returns the same value", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer inst.Release()

				Expect(inst.Value).To(Equal(1))
			})

			It("calls the factory's releaser only once all releasers have been called", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = firstInstance.Release()
				firstInstance.Releaser = nil
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				err = inst.Release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeFalse())
			})

			It("returns a new value if acquired after all releasers have been called", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = firstInstance.Release()
				firstInstance.Releaser = nil
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				err = inst.Release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeFalse())

				inst, err = policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer inst.Release()

				Expect(inst.Value).To(Equal(2))
			})

			Describe("func Close()", func() {
				It("does nothing", func() {
					Expect(exists).To(BeTrue())
					err := policy.Close()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(exists).To(BeTrue())
				})
			})
		})
	})
})
