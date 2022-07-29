package ferrite

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Int returns the value of an environment variable as an int.
func Int(name string) Integer[int] {
	return Integer[int]{name: name}
}

// Int8 returns the value of an environment variable as an int8.
func Int8(name string) Integer[int8] {
	return Integer[int8]{name: name}
}

// Int16 returns the value of an environment variable as an int16.
func Int16(name string) Integer[int16] {
	return Integer[int16]{name: name}
}

// Int32 returns the value of an environment variable as an int32.
func Int32(name string) Integer[int32] {
	return Integer[int32]{name: name}
}

// Int64 returns the value of an environment variable as an int64.
func Int64(name string) Integer[int64] {
	return Integer[int64]{name: name}
}

// Integer parses and validates signed integer environment variables.
type Integer[T constraints.Integer] struct {
	name     string
	min, max *T
}

func (n Integer[T]) Min(min T) Integer[T] {
	n.min = &min
	return n
}

func (n Integer[T]) Max(max T) Integer[T] {
	n.max = &max
	return n
}

func (n Integer[T]) Get() (T, error) {
	s := os.Getenv(n.name)
	if s == "" {
		return 0, fmt.Errorf(
			"%s is empty, expected %s",
			n.name,
			n.expected(),
		)
	}

	v64, err := strconv.ParseInt(s, 10, bits[T]())
	if err != nil {
		if errors.Is(err, strconv.ErrRange) {
			return 0, fmt.Errorf(
				"%s is out of range, expected %s, got '%s'",
				n.name,
				n.expected(),
				s,
			)
		}

		return 0, fmt.Errorf(
			"%s is invalid, expected %s, got '%s'",
			n.name,
			n.expected(),
			s,
		)
	}

	v := T(v64)

	if n.min != nil && v < *n.min {
		return 0, fmt.Errorf(
			"%s is too low, expected %s, got %+d",
			n.name,
			n.expected(),
			v,
		)
	}

	if n.max != nil && v > *n.max {
		return 0, fmt.Errorf(
			"%s is too high, expected %s, got %+d",
			n.name,
			n.expected(),
			v,
		)
	}

	return T(v64), nil
}

func (n Integer[T]) expected() string {
	if n.min != nil {
		if n.max != nil {
			return fmt.Sprintf("a value between %+d and %+d", *n.min, *n.max)
		}

		return fmt.Sprintf("%+d or greater", *n.min)
	}

	if n.max != nil {
		return fmt.Sprintf("%+d or lower", *n.max)
	}

	return fmt.Sprintf("%d-bit signed integer", bits[T]())
}
