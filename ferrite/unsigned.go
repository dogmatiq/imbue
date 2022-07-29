package ferrite

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// Unsigned returns the value of an environment variable as an unsigned integer.
func Unsigned[T constraints.Unsigned](name string) IntVar[T, uint64] {
	return IntVar[T, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}
