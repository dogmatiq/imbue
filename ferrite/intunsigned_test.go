package ferrite_test

import (
	"os"

	. "github.com/dogmatiq/imbue/ferrite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func UintX()", func() {
	When("the environment variable contains a valid integer", func() {
		DescribeTable(
			"it returns the integer value",
			func(value string, expect int) {
				os.Setenv("FERRITE_VALUE", value)

				// Uint
				{
					actual, err := Uint("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(uint(expect)))
				}

				// Uint8
				{
					actual, err := Uint8("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(uint8(expect)))
				}

				// Uint16
				{
					actual, err := Uint16("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(uint16(expect)))
				}

				// Uint32
				{
					actual, err := Uint32("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(uint32(expect)))
				}

				// Uint64
				{
					actual, err := Uint64("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(uint64(expect)))
				}
			},
			Entry(
				"zero",
				"0",
				0,
			),
			Entry(
				"positive",
				"100",
				100,
			),
		)
	})

	When("using the Min() constraint", func() {
		When("value >= minimum", func() {
			DescribeTable(
				"returns the value",
				func(value string, min, expect int) {
					os.Setenv("FERRITE_VALUE", value)

					// Uint
					{
						actual, err := Uint("FERRITE_VALUE").Min(uint(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint(expect)))
					}

					// Uint8
					{
						actual, err := Uint8("FERRITE_VALUE").Min(uint8(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint8(expect)))
					}

					// Uint16
					{
						actual, err := Uint16("FERRITE_VALUE").Min(uint16(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint16(expect)))
					}

					// Uint32
					{
						actual, err := Uint32("FERRITE_VALUE").Min(uint32(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint32(expect)))
					}

					// Uint64
					{
						actual, err := Uint64("FERRITE_VALUE").Min(uint64(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint64(expect)))
					}
				},
				Entry(
					"value == minimum",
					"50",
					50,
					50,
				),
				Entry(
					"value > minimum",
					"100",
					50,
					100,
				),
			)
		})

		When("value < minimum", func() {
			DescribeTable(
				"it returns an error",
				func(value string, min int, expect string) {
					os.Setenv("FERRITE_VALUE", value)

					// Uint
					{
						_, err := Uint("FERRITE_VALUE").Min(uint(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint8
					{
						_, err := Uint8("FERRITE_VALUE").Min(uint8(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint16
					{
						_, err := Uint16("FERRITE_VALUE").Min(uint16(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint32
					{
						_, err := Uint32("FERRITE_VALUE").Min(uint32(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint64
					{
						_, err := Uint64("FERRITE_VALUE").Min(uint64(min)).Get()
						Expect(err).To(MatchError(expect))
					}
				},
				Entry(
					"value < minimum",
					"50",
					100,
					"FERRITE_VALUE is too low, expected +100 or greater, got +50",
				),
			)
		})
	})

	When("using the Max() constraint", func() {
		When("value <= maximum", func() {
			DescribeTable(
				"returns the value",
				func(value string, max, expect int) {
					os.Setenv("FERRITE_VALUE", value)

					// Uint
					{
						actual, err := Uint("FERRITE_VALUE").Max(uint(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint(expect)))
					}

					// Uint8
					{
						actual, err := Uint8("FERRITE_VALUE").Max(uint8(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint8(expect)))
					}

					// Uint16
					{
						actual, err := Uint16("FERRITE_VALUE").Max(uint16(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint16(expect)))
					}

					// Uint32
					{
						actual, err := Uint32("FERRITE_VALUE").Max(uint32(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint32(expect)))
					}

					// Uint64
					{
						actual, err := Uint64("FERRITE_VALUE").Max(uint64(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(uint64(expect)))
					}
				},
				Entry(
					"value == maximum",
					"50",
					50,
					50,
				),
				Entry(
					"value < maximum",
					"50",
					100,
					50,
				),
			)
		})

		When("value > maximum", func() {
			DescribeTable(
				"it returns an error",
				func(value string, max int, expect string) {
					os.Setenv("FERRITE_VALUE", value)

					// Uint
					{
						_, err := Uint("FERRITE_VALUE").Max(uint(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint8
					{
						_, err := Uint8("FERRITE_VALUE").Max(uint8(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint16
					{
						_, err := Uint16("FERRITE_VALUE").Max(uint16(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint32
					{
						_, err := Uint32("FERRITE_VALUE").Max(uint32(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Uint64
					{
						_, err := Uint64("FERRITE_VALUE").Max(uint64(max)).Get()
						Expect(err).To(MatchError(expect))
					}
				},
				Entry(
					"value > maximum",
					"100",
					50,
					"FERRITE_VALUE is too high, expected +50 or lower, got +100",
				),
			)
		})
	})

	When("using the Min() and Max() constraints together", func() {
		It("explains the range in error messages", func() {
			os.Setenv("FERRITE_VALUE", "3")
			_, err := Uint("FERRITE_VALUE").Min(1).Max(2).Get()
			Expect(err).To(MatchError("FERRITE_VALUE is too high, expected a value between +1 and +2, got +3"))
		})
	})

	When("the environment variable does not contain a valid integer", func() {
		DescribeTable(
			"it returns an error",
			func(value string, expect string) {
				os.Setenv("FERRITE_VALUE", value)

				// Uint
				{
					_, err := Uint("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Uint8
				{
					_, err := Uint8("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Uint16
				{
					_, err := Uint16("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Uint32
				{
					_, err := Uint32("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Uint64
				{
					_, err := Uint64("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}
			},
			Entry(
				"empty",
				"",
				`FERRITE_VALUE is empty, expected (\d+)-bit unsigned integer`,
			),
			Entry(
				"non-numeric",
				"<invalid>",
				`FERRITE_VALUE is invalid, expected (\d+)-bit unsigned integer, got '<invalid>'`,
			),
			Entry(
				"decimal",
				"5.1",
				`FERRITE_VALUE is invalid, expected (\d+)-bit unsigned integer, got '5.1'`,
			),
			Entry(
				"underflow",
				"-100000000000000000000000000000000000",
				`FERRITE_VALUE is invalid, expected (\d+)-bit unsigned integer, got '-100000000000000000000000000000000000'`,
			),
			Entry(
				"overflow",
				"100000000000000000000000000000000000",
				`FERRITE_VALUE is out of range, expected (\d+)-bit unsigned integer, got '100000000000000000000000000000000000'`,
			),
		)
	})
})
