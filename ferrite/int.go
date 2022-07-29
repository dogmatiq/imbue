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

func (v IntVar[T, L]) Name() string {
	return v.name
}

// Min sets a minimum acceptable value.
func (v IntVar[T, L]) Min(min T) IntVar[T, L] {
	v.min = &min
	return v
}

// Max sets a maximum acceptable value.
func (v IntVar[T, L]) Max(max T) IntVar[T, L] {
	v.max = &max
	return v
}

// Default sets a default value to use when the environment variable is not
// defined.
func (v IntVar[T, L]) Default(def T) IntVar[T, L] {
	v.def = &def
	return v
}

// Get returns the integer value.
//
// It panics if the environment variable is not a valid integer or does not meet
// the constraints..
func (v IntVar[T, L]) Value() T {
	value, err := v.Parse()
	if err != nil {
		panic(err)
	}

	return value
}

// Parse returns the integer value.
//
// It returns an error if the environment variable is invalid or does not meet
// the min/max constraints.
func (v IntVar[T, L]) Parse() (T, error) {
	s := os.Getenv(v.name)
	if s == "" {
		if v.def != nil {
			return *v.def, nil
		}

		return 0, fmt.Errorf(
			"%s is empty, expected %s",
			v.name,
			v.expected(),
		)
	}

	value64, err := v.parse(s, 10, bits[T]())
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			return 0, fmt.Errorf(
				"%s is out of range, expected %s, got %s",
				v.name,
				v.expected(),
				s,
			)
		}

		return 0, fmt.Errorf(
			"%s is invalid, expected %s, got %q",
			v.name,
			v.expected(),
			s,
		)
	}

	value := T(value64)

	if v.min != nil && value < *v.min {
		return 0, fmt.Errorf(
			"%s is too low, expected %s, got %+d",
			v.name,
			v.expected(),
			value,
		)
	}

	if v.max != nil && value > *v.max {
		return 0, fmt.Errorf(
			"%s is too high, expected %s, got %+d",
			v.name,
			v.expected(),
			value,
		)
	}

	return T(value64), nil
}

// expected returns a description of the expected value.
func (v IntVar[T, L]) expected() string {
	if v.min != nil {
		if v.max != nil {
			return fmt.Sprintf("a value between %+d and %+d", *v.min, *v.max)
		}

		return fmt.Sprintf("%+d or greater", *v.min)
	}

	if v.max != nil {
		return fmt.Sprintf("%+d or lower", *v.max)
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
