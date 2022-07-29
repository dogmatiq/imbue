package ferrite_test

import (
	"os"

	. "github.com/dogmatiq/imbue/ferrite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Int()", func() {
	When("the environment variable contains a valid integer", func() {
		DescribeTable(
			"it returns the integer value",
			func(value string, expect int) {
				os.Setenv("FERRITE_VALUE", value)

				// Int
				{
					actual, err := Int("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(expect))
				}

				// Int8
				{
					actual, err := Int8("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(int8(expect)))
				}

				// Int16
				{
					actual, err := Int16("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(int16(expect)))
				}

				// Int32
				{
					actual, err := Int32("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(int32(expect)))
				}

				// Int64
				{
					actual, err := Int64("FERRITE_VALUE").Get()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(actual).To(Equal(int64(expect)))
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
			Entry(
				"negative",
				"-100",
				-100,
			),
		)
	})

	When("using the Min() constraint", func() {
		When("value >= minimum", func() {
			DescribeTable(
				"returns the value",
				func(value string, min, expect int) {
					os.Setenv("FERRITE_VALUE", value)

					// Int
					{
						actual, err := Int("FERRITE_VALUE").Min(min).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(expect))
					}

					// Int8
					{
						actual, err := Int8("FERRITE_VALUE").Min(int8(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int8(expect)))
					}

					// Int16
					{
						actual, err := Int16("FERRITE_VALUE").Min(int16(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int16(expect)))
					}

					// Int32
					{
						actual, err := Int32("FERRITE_VALUE").Min(int32(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int32(expect)))
					}

					// Int64
					{
						actual, err := Int64("FERRITE_VALUE").Min(int64(min)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int64(expect)))
					}
				},
				Entry(
					"positive value == positive minimum",
					"50",
					50,
					50,
				),
				Entry(
					"positive value > positive minimum",
					"100",
					50,
					100,
				),
				Entry(
					"positive value > negative minimum",
					"50",
					-50,
					50,
				),
				Entry(
					"negative value == negative minimum",
					"-50",
					-50,
					-50,
				),
				Entry(
					"negative value > negative minimum",
					"-50",
					-100,
					-50,
				),
			)
		})

		When("value < minimum", func() {
			DescribeTable(
				"it returns an error",
				func(value string, min int, expect string) {
					os.Setenv("FERRITE_VALUE", value)

					// Int
					{
						_, err := Int("FERRITE_VALUE").Min(min).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int8
					{
						_, err := Int8("FERRITE_VALUE").Min(int8(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int16
					{
						_, err := Int16("FERRITE_VALUE").Min(int16(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int32
					{
						_, err := Int32("FERRITE_VALUE").Min(int32(min)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int64
					{
						_, err := Int64("FERRITE_VALUE").Min(int64(min)).Get()
						Expect(err).To(MatchError(expect))
					}
				},
				Entry(
					"positive value < positive minimum",
					"50",
					100,
					"FERRITE_VALUE is too low, expected +100 or greater, got +50",
				),
				Entry(
					"negative value < positive minimum",
					"-50",
					100,
					"FERRITE_VALUE is too low, expected +100 or greater, got -50",
				),
				Entry(
					"negative value < negative minimum",
					"-100",
					-50,
					"FERRITE_VALUE is too low, expected -50 or greater, got -100",
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

					// Int
					{
						actual, err := Int("FERRITE_VALUE").Max(max).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(expect))
					}

					// Int8
					{
						actual, err := Int8("FERRITE_VALUE").Max(int8(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int8(expect)))
					}

					// Int16
					{
						actual, err := Int16("FERRITE_VALUE").Max(int16(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int16(expect)))
					}

					// Int32
					{
						actual, err := Int32("FERRITE_VALUE").Max(int32(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int32(expect)))
					}

					// Int64
					{
						actual, err := Int64("FERRITE_VALUE").Max(int64(max)).Get()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(actual).To(Equal(int64(expect)))
					}
				},
				Entry(
					"positive value == positive maximum",
					"50",
					50,
					50,
				),
				Entry(
					"positive value < positive maximum",
					"50",
					100,
					50,
				),
				Entry(
					"negative value < positive maximum",
					"-50",
					50,
					-50,
				),
				Entry(
					"negative value == negative maximum",
					"-50",
					-50,
					-50,
				),
				Entry(
					"negative value < negative maximum",
					"-100",
					-50,
					-100,
				),
			)
		})

		When("value > maximum", func() {
			DescribeTable(
				"it returns an error",
				func(value string, max int, expect string) {
					os.Setenv("FERRITE_VALUE", value)

					// Int
					{
						_, err := Int("FERRITE_VALUE").Max(max).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int8
					{
						_, err := Int8("FERRITE_VALUE").Max(int8(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int16
					{
						_, err := Int16("FERRITE_VALUE").Max(int16(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int32
					{
						_, err := Int32("FERRITE_VALUE").Max(int32(max)).Get()
						Expect(err).To(MatchError(expect))
					}

					// Int64
					{
						_, err := Int64("FERRITE_VALUE").Max(int64(max)).Get()
						Expect(err).To(MatchError(expect))
					}
				},
				Entry(
					"positive value > positive maximum",
					"100",
					50,
					"FERRITE_VALUE is too high, expected +50 or lower, got +100",
				),
				Entry(
					"positive value > negative minimum",
					"100",
					-50,
					"FERRITE_VALUE is too high, expected -50 or lower, got +100",
				),
				Entry(
					"negative value > negative maximum",
					"-50",
					-100,
					"FERRITE_VALUE is too high, expected -100 or lower, got -50",
				),
			)
		})
	})

	When("using the Min() and Max() constraints together", func() {
		It("explains the range in error messages", func() {
			os.Setenv("FERRITE_VALUE", "2")
			_, err := Int("FERRITE_VALUE").Min(-1).Max(+1).Get()
			Expect(err).To(MatchError("FERRITE_VALUE is too high, expected a value between -1 and +1, got +2"))
		})
	})

	When("using the Default() option", func() {
		When("the environment variable is undefined", func() {
			It("returns the default value", func() {
				os.Unsetenv("FERRITE_VALUE")

				actual, err := Int("FERRITE_VALUE").Default(10).Get()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual).To(Equal(10))
			})
		})

		When("the environment variable is empty", func() {
			It("returns the default value", func() {
				os.Setenv("FERRITE_VALUE", "")

				actual, err := Int("FERRITE_VALUE").Default(10).Get()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(actual).To(Equal(10))
			})
		})
	})

	When("the environment variable does not contain a valid integer", func() {
		DescribeTable(
			"it returns an error",
			func(value string, expect string) {
				os.Setenv("FERRITE_VALUE", value)

				// Int
				{
					_, err := Int("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Int8
				{
					_, err := Int8("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Int16
				{
					_, err := Int16("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Int32
				{
					_, err := Int32("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}

				// Int64
				{
					_, err := Int64("FERRITE_VALUE").Get()
					Expect(err).To(MatchError(MatchRegexp(expect)))
				}
			},
			Entry(
				"empty",
				"",
				`FERRITE_VALUE is empty, expected (\d+)-bit signed integer`,
			),
			Entry(
				"non-numeric",
				"<invalid>",
				`FERRITE_VALUE is invalid, expected (\d+)-bit signed integer, got '<invalid>'`,
			),
			Entry(
				"decimal",
				"5.1",
				`FERRITE_VALUE is invalid, expected (\d+)-bit signed integer, got '5.1'`,
			),
			Entry(
				"underflow",
				"-100000000000000000000000000000000000",
				`FERRITE_VALUE is out of range, expected (\d+)-bit signed integer, got '-100000000000000000000000000000000000'`,
			),
			Entry(
				"overflow",
				"100000000000000000000000000000000000",
				`FERRITE_VALUE is out of range, expected (\d+)-bit signed integer, got '100000000000000000000000000000000000'`,
			),
		)
	})
})
