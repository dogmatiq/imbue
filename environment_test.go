package imbue_test

import (
	"context"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type (
	envTestString   imbue.EnvironmentVariable[string]
	envTestBytes    imbue.EnvironmentVariable[[]byte]
	envTestBool     imbue.EnvironmentVariable[bool]
	envTestDuration imbue.EnvironmentVariable[time.Duration]

	envTestInt   imbue.EnvironmentVariable[int]
	envTestInt16 imbue.EnvironmentVariable[int16]
	envTestInt32 imbue.EnvironmentVariable[int32]
	envTestInt64 imbue.EnvironmentVariable[int64]

	envTestUint   imbue.EnvironmentVariable[uint]
	envTestUint16 imbue.EnvironmentVariable[uint16]
	envTestUint32 imbue.EnvironmentVariable[uint32]
	envTestUint64 imbue.EnvironmentVariable[uint64]

	envTestFloat32 imbue.EnvironmentVariable[float32]
	envTestFloat64 imbue.EnvironmentVariable[float64]
)

var _ = Describe("func FromEnvironment()", func() {
	It("can parse environment variables", func() {
		expectEnv[envTestString]("ENV_TEST_STRING", "<value>", "<value>")
		expectEnv[envTestBytes]("ENV_TEST_BYTES", "<value>", []byte("<value>"))
		expectEnv[envTestDuration]("ENV_TEST_DURATION", "1.5s", 1500*time.Millisecond)

		expectEnv[envTestBool]("ENV_TEST_BOOL", "TrUe", true)
		expectEnv[envTestBool]("ENV_TEST_BOOL", "fAlSe", false)
		expectEnv[envTestBool]("ENV_TEST_BOOL", "YeS", true)
		expectEnv[envTestBool]("ENV_TEST_BOOL", "nO", false)
		expectEnv[envTestBool]("ENV_TEST_BOOL", "On", true)
		expectEnv[envTestBool]("ENV_TEST_BOOL", "oFf", false)

		expectEnv[envTestInt]("ENV_TEST_INT", "-123", int(-123))
		expectEnv[envTestInt16]("ENV_TEST_INT16", "-123", int16(-123))
		expectEnv[envTestInt32]("ENV_TEST_INT32", "-123", int32(-123))
		expectEnv[envTestInt64]("ENV_TEST_INT64", "-123", int64(-123))

		expectEnv[envTestUint]("ENV_TEST_UINT", "123", uint(123))
		expectEnv[envTestUint16]("ENV_TEST_UINT16", "123", uint16(123))
		expectEnv[envTestUint32]("ENV_TEST_UINT32", "123", uint32(123))
		expectEnv[envTestUint64]("ENV_TEST_UINT64", "123", uint64(123))

		expectEnv[envTestFloat32]("ENV_TEST_FLOAT32", "-123.45", float32(-123.45))
		expectEnv[envTestFloat64]("ENV_TEST_FLOAT64", "-123.45", float64(-123.45))
	})

	It("returns an error when a bool cannot be parsed", func() {
		expectEnvError[envTestBool, bool](
			"ENV_TEST_BOOL",
			"<not-bool>",
			`the ENV_TEST_BOOL environment variable ("<not-bool>") is invalid: expected one of "true", "false", "yes", "no", "on" or "off"`,
		)
	})

	It("returns an error when a duration cannot be parsed", func() {
		expectEnvError[envTestDuration, time.Duration](
			"ENV_TEST_DURATION",
			"<not-duration>",
			`the ENV_TEST_DURATION environment variable ("<not-duration>") is invalid: expected duration (e.g. "300ms", "-1.5h" or "2h45m")`,
		)
	})

	It("returns an error when an int cannot be parsed", func() {
		expectEnvError[envTestInt, int](
			"ENV_TEST_INT",
			"<not-numeric>",
			fmt.Sprintf(
				`the ENV_TEST_INT environment variable ("<not-numeric>") is invalid: expected integer between %d and %d`,
				math.MinInt,
				math.MaxInt,
			),
		)
	})

	It("returns an error when an int16 cannot be parsed", func() {
		expectEnvError[envTestInt16, int16](
			"ENV_TEST_INT16",
			"<not-numeric>",
			`the ENV_TEST_INT16 environment variable ("<not-numeric>") is invalid: expected integer between -32768 and 32767`,
		)
	})

	It("returns an error when an int32 cannot be parsed", func() {
		expectEnvError[envTestInt32, int32](
			"ENV_TEST_INT32",
			"<not-numeric>",
			`the ENV_TEST_INT32 environment variable ("<not-numeric>") is invalid: expected integer between -2147483648 and 2147483647`,
		)
	})

	It("returns an error when an int64 cannot be parsed", func() {
		expectEnvError[envTestInt64, int64](
			"ENV_TEST_INT64",
			"<not-numeric>",
			`the ENV_TEST_INT64 environment variable ("<not-numeric>") is invalid: expected integer between -9223372036854775808 and 9223372036854775807`,
		)
	})

	It("returns an error when a uint cannot be parsed", func() {
		expectEnvError[envTestUint, uint](
			"ENV_TEST_UINT",
			"<not-numeric>",
			`the ENV_TEST_UINT environment variable ("<not-numeric>") is invalid: expected integer between 0 and 18446744073709551615`,
		)
	})

	It("returns an error when a uint16 cannot be parsed", func() {
		expectEnvError[envTestUint16, uint16](
			"ENV_TEST_UINT16",
			"<not-numeric>",
			`the ENV_TEST_UINT16 environment variable ("<not-numeric>") is invalid: expected integer between 0 and 65535`,
		)
	})

	It("returns an error when a uint32 cannot be parsed", func() {
		expectEnvError[envTestUint32, uint32](
			"ENV_TEST_UINT32",
			"<not-numeric>",
			`the ENV_TEST_UINT32 environment variable ("<not-numeric>") is invalid: expected integer between 0 and 4294967295`,
		)
	})

	It("returns an error when a uint64 cannot be parsed", func() {
		expectEnvError[envTestUint64, uint64](
			"ENV_TEST_UINT64",
			"<not-numeric>",
			`the ENV_TEST_UINT64 environment variable ("<not-numeric>") is invalid: expected integer between 0 and 18446744073709551615`,
		)
	})

	It("returns an error when a float32 cannot be parsed", func() {
		expectEnvError[envTestFloat32, float32](
			"ENV_TEST_FLOAT32",
			"<not-numeric>",
			`the ENV_TEST_FLOAT32 environment variable ("<not-numeric>") is invalid: expected floating-point number`,
		)
	})

	It("returns an error when a float64 cannot be parsed", func() {
		expectEnvError[envTestFloat64, float64](
			"ENV_TEST_FLOAT64",
			"<not-numeric>",
			`the ENV_TEST_FLOAT64 environment variable ("<not-numeric>") is invalid: expected floating-point number`,
		)
	})

	It("returns an error when the environment variable is undefined", func() {
		con := imbue.New()
		defer con.Close()

		err := imbue.Invoke1(
			context.Background(),
			con,
			func(
				ctx context.Context,
				v imbue.FromEnvironment[envTestString, string],
			) error {
				panic("unexpected call")
			},
		)
		Expect(err).To(MatchError("the ENV_TEST_STRING environment variable is not defined"))
	})

	It("returns an error when the environment variable is empty", func() {
		os.Setenv("ENV_TEST_STRING", "")
		defer os.Unsetenv("ENV_TEST_STRING")

		con := imbue.New()
		defer con.Close()

		err := imbue.Invoke1(
			context.Background(),
			con,
			func(
				ctx context.Context,
				v imbue.FromEnvironment[envTestString, string],
			) error {
				panic("unexpected call")
			},
		)
		Expect(err).To(MatchError("the ENV_TEST_STRING environment variable is defined, but it is empty"))
	})

	It("panics if a constructor is declared for an optional type", func() {
		Expect(func() {
			con := imbue.New()
			defer con.Close()

			imbue.With0(
				con,
				func(
					ctx imbue.Context,
				) (imbue.FromEnvironment[envTestString, string], error) {
					panic("unexpected call")
				},
			)
		}).To(
			PanicWith(
				MatchRegexp(
					`explicit declaration of imbue\.FromEnvironment\[imbue.envTestString,string\] constructor \(environment_test\.go:\d+\) is disallowed`,
				),
			),
		)
	})
})

func expectEnv[
	V imbue.EnvironmentVariable[T],
	T imbue.Parseable,
](
	name string,
	raw string,
	value T,
) {
	os.Setenv(name, raw)
	defer os.Unsetenv(name)

	con := imbue.New()
	defer con.Close()

	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			v imbue.FromEnvironment[V, T],
		) error {
			Expect(v.Name()).To(Equal(name))
			Expect(v.Value()).To(Equal(value))
			Expect(v.String()).To(Equal(raw))
			return nil
		},
	)
	Expect(err).ShouldNot(HaveOccurred())
}

func expectEnvError[
	V imbue.EnvironmentVariable[T],
	T imbue.Parseable,
](name, raw, message string) {
	os.Setenv(name, raw)
	defer os.Unsetenv(name)

	con := imbue.New()
	defer con.Close()

	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			v imbue.FromEnvironment[V, T],
		) error {
			panic("unexpected call")
		},
	)
	Expect(err).To(MatchError(message))
}
