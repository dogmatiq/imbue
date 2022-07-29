package ferrite

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Int returns the value of an environment variable as an int.
func Int(name string) Integer[int, int64] {
	return Integer[int, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}

// Int8 returns the value of an environment variable as an int8.
func Int8(name string) Integer[int8, int64] {
	return Integer[int8, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}

// Int16 returns the value of an environment variable as an int16.
func Int16(name string) Integer[int16, int64] {
	return Integer[int16, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}

// Int32 returns the value of an environment variable as an int32.
func Int32(name string) Integer[int32, int64] {
	return Integer[int32, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}

// Int64 returns the value of an environment variable as an int64.
func Int64(name string) Integer[int64, int64] {
	return Integer[int64, int64]{
		name:  name,
		parse: strconv.ParseInt,
	}
}

// Uint returns the value of an environment variable as a uint.
func Uint(name string) Integer[uint, uint64] {
	return Integer[uint, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}

// Uint8 returns the value of an environment variable as a uint8.
func Uint8(name string) Integer[uint8, uint64] {
	return Integer[uint8, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}

// Uint16 returns the value of an environment variable as a uint16.
func Uint16(name string) Integer[uint16, uint64] {
	return Integer[uint16, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}

// Uint32 returns the value of an environment variable as a uint32.
func Uint32(name string) Integer[uint32, uint64] {
	return Integer[uint32, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}

// Uint64 returns the value of an environment variable as a uint64.
func Uint64(name string) Integer[uint64, uint64] {
	return Integer[uint64, uint64]{
		name:  name,
		parse: strconv.ParseUint,
	}
}

// Integer parses and validates signed integer environment variables.
type Integer[T, B constraints.Integer] struct {
	name     string
	parse    func(string, int, int) (B, error)
	min, max *T
}

// Min sets a minimum acceptable value.
func (i Integer[T, B]) Min(min T) Integer[T, B] {
	i.min = &min
	return i
}

// Max sets a maximum acceptable value.
func (i Integer[T, B]) Max(max T) Integer[T, B] {
	i.max = &max
	return i
}

// Get returns the integer value.
//
// It returns an error if the environment variable is invalid or does not meet
// the min/max constraints.
func (i Integer[T, B]) Get() (T, error) {
	s := os.Getenv(i.name)
	if s == "" {
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
func (i Integer[T, B]) expected() string {
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

func isSigned[T constraints.Integer]() bool {
	var zero T
	return (zero - 1) < 0
}
