package ferrite

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// Unsigned returns the value of an environment variable as an unsigned integer.
func Unsigned[T constraints.Unsigned](name string) Int[T, uint64] {
	return Int[T, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}
