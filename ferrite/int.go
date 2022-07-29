package ferrite

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// IntVar parses and validates integer environment variables.
//
// T is the type of integer to produce. L is the largest integer type that can
// represent values of type T.
type IntVar[T, L constraints.Integer] struct {
	name          string
	parse         func(string, int, int) (L, error)
	min, max, def *T
}

func (i IntVar[T, L]) Name() string {
	return i.name
}

// Min sets a minimum acceptable value.
func (i IntVar[T, L]) Min(min T) IntVar[T, L] {
	i.min = &min
	return i
}

// Max sets a maximum acceptable value.
func (i IntVar[T, L]) Max(max T) IntVar[T, L] {
	i.max = &max
	return i
}

// Default sets a default value to use when the environment variable is not
// defined.
func (i IntVar[T, L]) Default(def T) IntVar[T, L] {
	i.def = &def
	return i
}

// Get returns the integer value.
//
// It panics if the environment variable is not a valid integer or does not meet
// the constraints..
func (i IntVar[T, L]) Value() T {
	v, err := i.Parse()
	if err != nil {
		panic(err)
	}

	return v
}

// Parse returns the integer value.
//
// It returns an error if the environment variable is invalid or does not meet
// the min/max constraints.
func (i IntVar[T, L]) Parse() (T, error) {
	s := os.Getenv(i.name)
	if s == "" {
		if i.def != nil {
			return *i.def, nil
		}

		return 0, fmt.Errorf(
			"%s is empty, expected %s",
			i.name,
			i.expected(),
		)
	}

	v64, err := i.parse(s, 10, bits[T]())
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			return 0, fmt.Errorf(
				"%s is out of range, expected %s, got '%s'",
				i.name,
				i.expected(),
				s,
			)
		}

		return 0, fmt.Errorf(
			"%s is invalid, expected %s, got '%s'",
			i.name,
			i.expected(),
			s,
		)
	}

	v := T(v64)

	if i.min != nil && v < *i.min {
		return 0, fmt.Errorf(
			"%s is too low, expected %s, got %+d",
			i.name,
			i.expected(),
			v,
		)
	}

	if i.max != nil && v > *i.max {
		return 0, fmt.Errorf(
			"%s is too high, expected %s, got %+d",
			i.name,
			i.expected(),
			v,
		)
	}

	return T(v64), nil
}

// expected returns a description of the expected value.
func (i IntVar[T, L]) expected() string {
	if i.min != nil {
		if i.max != nil {
			return fmt.Sprintf("a value between %+d and %+d", *i.min, *i.max)
		}

		return fmt.Sprintf("%+d or greater", *i.min)
	}

	if i.max != nil {
		return fmt.Sprintf("%+d or lower", *i.max)
	}

	if isSigned[T]() {
		return fmt.Sprintf("%d-bit signed integer", bits[T]())
	}

	return fmt.Sprintf("%d-bit unsigned integer", bits[T]())
}

// isSigned returns true if T is a signed integer type.
func isSigned[T constraints.Integer]() bool {
	var zero T
	return (zero - 1) < 0
}

// bits returns the number of bits used to store a represent of type T.
func bits[T any]() int {
	var zero T
	return int(unsafe.Sizeof(zero)) * 8
}
