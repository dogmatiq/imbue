package imbue

import (
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

// EnvironmentVariable is a constraint for a type that identifies an environment
// variable.
//
// Environment variables are declared by declaring a type that uses
// imbue.EnvironmentVariable[T] as its underlying type.
//
// The name of the declared type is used as the name of the environment
// variable.
//
// T is the type that is produced by parsing the environment variable's value.
type EnvironmentVariable[T Parseable] interface {
	variableOfType(T)
}

// Parseable is a constraint that identifies the set of types that can be parsed
// from their string representation, such as in an environment variable.
type Parseable interface {
	string |
		int | int16 | int32 | int64 |
		uint | uint16 | uint32 | uint64 |
		float32 | float64
}

// FromEnvironment requests a dependency from the environment.
//
// It is used as a parameter type within user-defined functions passed to
// WithX(), DecorateX(), and InvokeX() to request a dependency of type T that is
// parsed from the environment variable V.
type FromEnvironment[V EnvironmentVariable[T], T Parseable] struct {
	value T
}

// Name returns the name of the environment variable.
func (v FromEnvironment[V, T]) Name() string {
	return typeOf[V]().Name()
}

// Value returns the parsed value of the environment variable.
func (v FromEnvironment[V, T]) Value() T {
	return v.value
}

// String returns the original string value of the environment variable.
func (v FromEnvironment[V, T]) String() string {
	return os.Getenv(v.Name())
}

// constructSelf constructs the environment variable in-place.
func (v *FromEnvironment[V, T]) constructSelf(ctx *Context) error {
	name := v.Name()
	value, ok := os.LookupEnv(name)
	if !ok {
		return fmt.Errorf(
			"the %s environment variable is not defined",
			name,
		)
	}

	if value == "" {
		return fmt.Errorf(
			"the %s environment variable is defined, but it is empty",
			name,
		)
	}

	if err := parseInto(value, &v.value); err != nil {
		return fmt.Errorf(
			"the %s environment variable cannot be parsed: %w",
			name,
			err,
		)
	}

	return nil
}

// parseInto parses value into *out.
func parseInto(value string, out any) error {
	switch out := out.(type) {
	case *string:
		*out = value

	case *int:
		return parseInt(value, bits.UintSize, out)
	case *int16:
		return parseInt(value, 16, out)
	case *int32:
		return parseInt(value, 32, out)
	case *int64:
		return parseInt(value, 64, out)

	case *uint:
		return parseUint(value, bits.UintSize, out)
	case *uint16:
		return parseUint(value, 16, out)
	case *uint32:
		return parseUint(value, 32, out)
	case *uint64:
		return parseUint(value, 64, out)

	case *float32:
		return parseFloat(value, 32, out)
	case *float64:
		return parseFloat(value, 64, out)

	default:
		panic(fmt.Sprintf(
			"%s implements the Parseable constraint, but is not handled by the parser",
			reflect.TypeOf(out).Elem(),
		))
	}

	return nil
}

// parseInt parses a signed integer and assigns it to *out.
func parseInt[T constraints.Signed](value string, size int, out *T) error {
	n, err := strconv.ParseInt(value, 10, size)
	if err != nil {
		return fmt.Errorf(
			"%#v is not a valid %s",
			value,
			typeOf[T](),
		)
	}

	*out = T(n)
	return nil
}

// parseUint parses an unsigned integer and assigns it to *out.
func parseUint[T constraints.Unsigned](value string, size int, out *T) error {
	n, err := strconv.ParseUint(value, 10, size)
	if err != nil {
		return fmt.Errorf(
			"%#v is not a valid %s",
			value,
			typeOf[T](),
		)
	}

	*out = T(n)
	return nil
}

// parseFloat parses a float and assigns it to *out.
func parseFloat[T constraints.Float](value string, size int, out *T) error {
	n, err := strconv.ParseFloat(value, size)
	if err != nil {
		return fmt.Errorf(
			"%#v is not a valid %s",
			value,
			typeOf[T](),
		)
	}

	*out = T(n)
	return nil
}
