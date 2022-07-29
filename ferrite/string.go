package ferrite

import (
	"fmt"
	"os"
)

// String returns the value of an environment variable as a string-like type.
func String(name string) StringVar[string] {
	return StringAs[string](name)
}

// StringAs returns the value of an environment variable as a string-like type.
func StringAs[T ~string](name string) StringVar[T] {
	return StringVar[T]{
		name: name,
	}
}

// StringVar parses and validates string-like environment variables.
//
// T is the type of value to produce.
type StringVar[T ~string] struct {
	name     string
	min, max *int
	def      *T
}

func (v StringVar[T]) Name() string {
	return v.name
}

// MinLen sets a minimum acceptable length (in bytes).
func (v StringVar[T]) MinLen(n int) StringVar[T] {
	v.min = &n
	return v
}

// Max sets a maximum acceptable length (in bytes).
func (v StringVar[T]) MaxLen(n int) StringVar[T] {
	v.max = &n
	return v
}

// Default sets a default value to use when the environment variable is not
// defined.
func (v StringVar[T]) Default(def T) StringVar[T] {
	v.def = &def
	return v
}

// Get returns the integer value.
//
// It panics if the environment variable is not a valid integer or does not meet
// the constraints..
func (v StringVar[T]) Value() T {
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
func (v StringVar[T]) Parse() (T, error) {
	value := os.Getenv(v.name)
	if value == "" {
		if v.def != nil {
			return *v.def, nil
		}

		return "", fmt.Errorf(
			"%s is empty, expected %s",
			v.name,
			v.expected(),
		)
	}

	if v.min != nil && len(value) < *v.min {
		return "", fmt.Errorf(
			"%s is too short, expected %s, got %q (length = %d)",
			v.name,
			v.expected(),
			value,
			len(value),
		)
	}

	if v.max != nil && len(value) > *v.max {
		return "", fmt.Errorf(
			"%s is too long, expected %s, got %q (length = %d)",
			v.name,
			v.expected(),
			value,
			len(value),
		)
	}

	return T(value), nil
}

// expected returns a description of the expected value.
func (v StringVar[T]) expected() string {
	if v.min != nil {
		if v.max != nil {
			return fmt.Sprintf("between %d and %d bytes", *v.min, *v.max)
		}

		return fmt.Sprintf("%d bytes or more", *v.min)
	}

	if v.max != nil {
		return fmt.Sprintf("%d bytes or less", *v.max)
	}

	return "non-empty string"
}
