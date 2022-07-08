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

		factory = func(ctx context.Context) (int, Releaser, error) {
			if exists {
				panic("unexpected call")
			}

			exists = true
			release := func() error {
				exists = false
				return nil
			}

			count++
			return count, release, nil
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
				value, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer release()

				Expect(value).To(Equal(1))
			})

			// It("returns a no-op releaser", func() {
			// 	_, release, err := policy.Acquire(context.Background(), factory)
			// 	Expect(err).ShouldNot(HaveOccurred())

			// 	err = release()
			// 	Expect(err).ShouldNot(HaveOccurred())
			// 	Expect(count).To(Equal(1))
			// })

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
		var releaseFirst Releaser

		BeforeEach(func() {
			var err error
			_, releaseFirst, err = policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			if releaseFirst != nil {
				err := releaseFirst()
				Expect(err).ShouldNot(HaveOccurred())
			}
		})

		Describe("func Acquire()", func() {
			It("returns the same value", func() {
				value, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer release()

				Expect(value).To(Equal(1))
			})

			It("calls the factory's releaser only once all releasers have been called", func() {
				_, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = releaseFirst()
				releaseFirst = nil
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				err = release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeFalse())
			})

			It("returns a new value if acquired after all releasers have been called", func() {
				_, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())

				err = releaseFirst()
				releaseFirst = nil
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeTrue())

				err = release()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(exists).To(BeFalse())

				value, release, err := policy.Acquire(context.Background(), factory)
				Expect(err).ShouldNot(HaveOccurred())
				defer release()

				Expect(value).To(Equal(2))
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
