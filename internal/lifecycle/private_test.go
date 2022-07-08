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

		factory = func(ctx context.Context) (int, Releaser, error) {
			count++

			release := func() error {
				count--
				return nil
			}

			return count, release, nil
		}

		count = 0
	})

	AfterEach(func() {
		err := policy.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("func Acquire()", func() {
		It("returns a new value on each call", func() {
			value, release, err := policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
			defer release()

			Expect(value).To(Equal(1))

			value, release, err = policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())
			defer release()

			Expect(value).To(Equal(2))
		})

		It("returns the factory's releaser", func() {
			_, release, err := policy.Acquire(context.Background(), factory)
			Expect(err).ShouldNot(HaveOccurred())

			err = release()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(count).To(Equal(0))
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
