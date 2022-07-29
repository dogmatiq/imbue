package ferrite_test

import (
	"os"

	. "github.com/dogmatiq/imbue/ferrite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func String()", func() {
	When("the environment variable contains a non-empty string", func() {
		It("returns the value", func() {
			expect := "<value>"
			os.Setenv("FERRITE_VALUE", expect)

			actual, err := String("FERRITE_VALUE").Parse()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(actual).To(Equal(expect))
		})
	})

	When("using the MinLen() constraint", func() {
		When("length >= minimum", func() {
			DescribeTable(
				"returns the value",
				func(value string, min int) {
					os.Setenv("FERRITE_VALUE", value)

					actual, err := String("FERRITE_VALUE").MinLen(min).Parse()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(value))
				},
				Entry(
					"length == minimum",
					"<value>",
					7,
				),
				Entry(
					"length > minimum",
					"<value>",
					6,
				),
			)
		})

		When("length < minimum", func() {
			It("returns an error", func() {
				os.Setenv("FERRITE_VALUE", "<value>")

				_, err := String("FERRITE_VALUE").MinLen(8).Parse()
				Expect(err).To(MatchError(
					`FERRITE_VALUE is too short, expected 8 bytes or more, got "<value>" (length = 7)`,
				))
			})
		})
	})

	When("using the MaxLen() constraint", func() {
		When("length <= maximum", func() {
			DescribeTable(
				"returns the value",
				func(value string, max int) {
					os.Setenv("FERRITE_VALUE", value)

					actual, err := String("FERRITE_VALUE").MaxLen(max).Parse()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(value))
				},
				Entry(
					"length == maximum",
					"<value>",
					7,
				),
				Entry(
					"length < maximum",
					"<value>",
					8,
				),
			)
		})

		When("value > maximum", func() {
			It("returns an error", func() {
				os.Setenv("FERRITE_VALUE", "<value>")

				_, err := String("FERRITE_VALUE").MaxLen(6).Parse()
				Expect(err).To(MatchError(
					`FERRITE_VALUE is too long, expected 6 bytes or less, got "<value>" (length = 7)`,
				))
			})
		})
	})

	When("using the MinLen() and MaxLen() constraints together", func() {
		It("explains the range in error messages", func() {
			os.Setenv("FERRITE_VALUE", "<value>")
			_, err := String("FERRITE_VALUE").MinLen(1).MaxLen(6).Parse()
			Expect(err).To(MatchError(
				`FERRITE_VALUE is too long, expected between 1 and 6 bytes, got "<value>" (length = 7)`,
			))
		})
	})

	When("using the Default() option", func() {
		When("the environment variable is undefined", func() {
			It("returns the default value", func() {
				os.Unsetenv("FERRITE_VALUE")

				actual, err := String("FERRITE_VALUE").Default("<default>").Parse()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual).To(Equal("<default>"))
			})
		})

		When("the environment variable is empty", func() {
			It("returns the default value", func() {
				os.Setenv("FERRITE_VALUE", "")

				actual, err := String("FERRITE_VALUE").Default("<default>").Parse()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual).To(Equal("<default>"))
			})
		})
	})

	When("the environment variable is empty", func() {
		It("it returns an error", func() {
			os.Setenv("FERRITE_VALUE", "")

			_, err := String("FERRITE_VALUE").Parse()
			Expect(err).To(MatchError("FERRITE_VALUE is empty, expected non-empty string"))
		})
	})
})
