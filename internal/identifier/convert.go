package identifier

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToScreamingSnakeCase converts a Go identifier to SCREAMING_SNAKE_CASE.
func ToScreamingSnakeCase(camel string) string {
	var (
		wordIsUpper bool
		wordLen     int
		result      strings.Builder
	)

	for pos, ch := range camel {
		if pos == 0 {
			wordIsUpper = unicode.IsUpper(ch)
		} else if wordIsUpper {
			if unicode.IsLower(ch) {
				wordIsUpper = false

				if wordLen > 1 {
					wordLen = 0

					res := result.String()
					result.Reset()

					last, size := utf8.DecodeLastRuneInString(res)
					result.WriteString(res[:len(res)-size])

					result.WriteRune('_')
					result.WriteRune(last)
				}
			}
		} else {
			if unicode.IsUpper(ch) {
				wordIsUpper = true
				wordLen = 0
				result.WriteRune('_')
			}
		}

		result.WriteRune(unicode.ToUpper(ch))
		wordLen++
	}

	return result.String()
}
