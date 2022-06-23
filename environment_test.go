package imbue_test

import (
	"context"
	"os"

	"github.com/dogmatiq/imbue"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type (
	IMBUE_STRING imbue.EnvironmentVariable[string]

	IMBUE_INT   imbue.EnvironmentVariable[int]
	IMBUE_INT16 imbue.EnvironmentVariable[int16]

	IMBUE_UINT   imbue.EnvironmentVariable[uint]
	IMBUE_UINT16 imbue.EnvironmentVariable[uint16]
)

var _ = Describe("func FromEnvironment()", func() {
	var container *imbue.Container

	BeforeEach(func() {
		container = imbue.New()
	})

	AfterEach(func() {
		container.Close()
	})

	It("can parse environment variables", func() {
		expectEnv[IMBUE_STRING](container, "IMBUE_STRING", "<value>", "<value>")

		expectEnv[IMBUE_INT](container, "IMBUE_INT", "-123", int(-123))
		expectEnv[IMBUE_INT16](container, "IMBUE_INT16", "-123", int16(-123))

		expectEnv[IMBUE_UINT](container, "IMBUE_UINT", "123", uint(123))
		expectEnv[IMBUE_UINT16](container, "IMBUE_UINT16", "123", uint16(123))
	})

	It("returns an error when an int cannot be parsed", func() {
		expectEnvError[IMBUE_INT, int](
			container,
			"IMBUE_INT",
			"<not-numeric>",
			`the IMBUE_INT environment variable cannot be parsed: "<not-numeric>" is not a valid int`,
		)
	})

	It("returns an error when an int16 cannot be parsed", func() {
		expectEnvError[IMBUE_INT16, int16](
			container,
			"IMBUE_INT16",
			"<not-numeric>",
			`the IMBUE_INT16 environment variable cannot be parsed: "<not-numeric>" is not a valid int16`,
		)
	})

	It("returns an error when a uint cannot be parsed", func() {
		expectEnvError[IMBUE_UINT, uint](
			container,
			"IMBUE_UINT",
			"<not-numeric>",
			`the IMBUE_UINT environment variable cannot be parsed: "<not-numeric>" is not a valid uint`,
		)
	})

	It("returns an error when a uint16 cannot be parsed", func() {
		expectEnvError[IMBUE_UINT16, uint16](
			container,
			"IMBUE_UINT16",
			"<not-numeric>",
			`the IMBUE_UINT16 environment variable cannot be parsed: "<not-numeric>" is not a valid uint16`,
		)
	})

	It("returns an error when the environment variable is undefined", func() {
		err := imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				v imbue.FromEnvironment[IMBUE_STRING, string],
			) error {
				Fail("unexpected call")
				return nil
			},
		)
		Expect(err).To(MatchError("the IMBUE_STRING environment variable is not defined"))
	})

	It("returns an error when the environment variable is empty", func() {
		os.Setenv("IMBUE_STRING", "")
		defer os.Unsetenv("IMBUE_STRING")

		err := imbue.Invoke1(
			context.Background(),
			container,
			func(
				ctx context.Context,
				v imbue.FromEnvironment[IMBUE_STRING, string],
			) error {
				Fail("unexpected call")
				return nil
			},
		)
		Expect(err).To(MatchError("the IMBUE_STRING environment variable is defined, but it is empty"))
	})
})

func expectEnv[
	V imbue.EnvironmentVariable[T],
	T imbue.Parseable,
](
	con *imbue.Container,
	name string,
	stringValue string,
	parsedValue T,
) {
	os.Setenv(name, stringValue)
	defer os.Unsetenv(name)

	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			v imbue.FromEnvironment[V, T],
		) error {
			Expect(v.Name()).To(Equal(name))
			Expect(v.Value()).To(Equal(parsedValue))
			Expect(v.String()).To(Equal(stringValue))
			return nil
		},
	)
	Expect(err).ShouldNot(HaveOccurred())
}

func expectEnvError[
	V imbue.EnvironmentVariable[T],
	T imbue.Parseable,
](
	con *imbue.Container,
	name string,
	stringValue string,
	errorMessage string,
) {
	os.Setenv(name, stringValue)

	err := imbue.Invoke1(
		context.Background(),
		con,
		func(
			ctx context.Context,
			v imbue.FromEnvironment[V, T],
		) error {
			Fail("unexpected call")
			return nil
		},
	)
	Expect(err).To(MatchError(errorMessage))
}
