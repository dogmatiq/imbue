package identifier_test

import (
	. "github.com/dogmatiq/imbue/internal/identifier"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("func ToScreamingSnakeCase()", func() {
	DescribeTable(
		"it converts a string to SCREAMING_SNAKE_CASE",
		func(camel, expect string) {
			Expect(ToScreamingSnakeCase(camel)).To(Equal(expect))
		},
		Entry("2-word identity", "FOO_BAR", "FOO_BAR"),
		Entry("2-word camel case", "fooBar", "FOO_BAR"),
		Entry("2-word pascal case", "FooBar", "FOO_BAR"),
		Entry("2-word initialism at start", "FOOBar", "FOO_BAR"),
		Entry("2-word initialism at end", "FooBAR", "FOO_BAR"),

		Entry("3-word identity", "FOO_BAR_SPAM", "FOO_BAR_SPAM"),
		Entry("3-word camel case", "fooBarSpam", "FOO_BAR_SPAM"),
		Entry("3-word pascal case", "FooBarSpam", "FOO_BAR_SPAM"),
		Entry("3-word initialism at start", "FOOBarSpam", "FOO_BAR_SPAM"),
		Entry("3-word initialism at middle", "FooBARSpam", "FOO_BAR_SPAM"),
		Entry("3-word initialism at end", "FooBarSPAM", "FOO_BAR_SPAM"),

		Entry("numbers", "FooBar123Spam", "FOO_BAR123_SPAM"),
		Entry("numbers with initialism", "FooBAR123Spam", "FOO_BAR123_SPAM"),
	)
})
