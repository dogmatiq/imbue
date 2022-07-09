package lifecycle_test

import (
	"context"
	"errors"

	. "github.com/dogmatiq/imbue/internal/lifecycle"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Singleton", func() {
	var (
		policy  Singleton[int]
		factory Factory[int]
		exists  bool
	)

	BeforeEach(func() {
		policy = Singleton[int]{}

		factory = func(ctx context.Context) (*Instance[int], error) {
			if exists {
				panic("unexpected call")
			}

			exists = true

			return &Instance[int]{
				Value: 123,
				Releaser: func() error {
					exists = false
					return nil
				},
			}, nil
		}

		exists = false
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

				Expect(inst.Value).To(Equal(123))
			})

			It("returns a no-op releaser", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = inst.Release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())
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
			err := firstInstance.Release()
			Expect(err).ShouldNot(HaveOccurred())
		})

		Describe("func Acquire()", func() {
			It("returns the same value", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer inst.Release()

				Expect(inst.Value).To(Equal(123))
			})

			It("returns a no-op releaser", func() {
				inst, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = inst.Release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())
			})

			Describe("func Close()", func() {
				It("calls the factory's releaser", func() {
					err := policy.Close()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(exists).To(BeFalse())
				})
			})
		})
	})
})
