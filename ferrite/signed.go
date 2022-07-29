package ferrite

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// Signed returns the value of an environment variable as a signed integer.
func Signed[T constraints.Signed](name string) Int[T, int64] {
	return Int[T, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}
