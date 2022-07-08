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

		factory = func(ctx context.Context) (int, Releaser, error) {
			if exists {
				panic("unexpected call")
			}

			exists = true
			release := func() error {
				exists = false
				return nil
			}

			return 123, release, nil
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
				value, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer release()

				Expect(value).To(Equal(123))
			})

			It("returns a no-op releaser", func() {
				_, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())
			})

			It("returns an error if the factory returns an error", func() {
				factory = func(ctx context.Context) (int, Releaser, error) {
					return 0, nil, errors.New("<error>")
				}

				_, _, err := policy.Acquire(context.Background(), factory)
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
		var release Releaser

		BeforeEach(func() {
			var err error
			_, release, err = policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			err := release()
			Expect(err).ShouldNot(HaveOccurred())
		})

		Describe("func Acquire()", func() {
			It("returns the same value", func() {
				value, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer release()

				Expect(value).To(Equal(123))
			})

			It("returns a no-op releaser", func() {
				_, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = release()
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
