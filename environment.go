package imbue

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/dogmatiq/imbue/internal/identifier"
	"golang.org/x/exp/constraints"
)

// EnvironmentVariable is a constraint for a type that identifies an environment
// variable.
//
// Environment variables are read by declaring a type that uses
// imbue.EnvironmentVariable[T] as its underlying type.
//
// The environment variable name is obtained by converting the name of the type
// to "screaming snake case". For example, a Go type named "fooBar" would map
// to the environment variable "FOO_BAR".
//
// T is the type that is produced by parsing the environment variable's value.
// See the Parseable constraint for information on how each type is parsed.
type EnvironmentVariable[T Parseable] interface {
	variableOfType(T)
}

// Parseable is a constraint that identifies the set of types that can be parsed
// from their string representation, such as in an environment variable.
type Parseable interface {
	string |

		// Booleans are required to be explicitly set to one of "true", "false",
		// "yes", "no", "on" or "off". The value is case-insensitive.
		bool |

		// Signed integers are parsed using strconv.ParseInt() with "base
		// detection", meaning that values are assumed to be in base-10 unless
		// prefixed with '0b', '0', '0o' or '0x'.
		int | int16 | int32 | int64 |

		// Unsigned integers are parsed using strconv.ParseUint() using the same
		// base-detection rules as signed integers.
		uint | uint16 | uint32 | uint64 |

		// Floating point numbers are parsed using strconv.ParseFloat().
		float32 | float64
}

// FromEnvironment requests a dependency from the environment.
//
// It is used as a parameter type within user-defined functions passed to
// WithX(), DecorateX(), and InvokeX() to request a dependency of type T that is
// parsed from the environment variable V.
type FromEnvironment[V EnvironmentVariable[T], T Parseable] struct {
	name  string
	raw   string
	value T
}

// Name returns the name of the environment variable.
func (v FromEnvironment[V, T]) Name() string {
	return v.name
}

// Value returns the parsed value of the environment variable.
func (v FromEnvironment[V, T]) Value() T {
	return v.value
}

// String returns the raw string value of the environment variable.
func (v FromEnvironment[V, T]) String() string {
	return v.raw
}

// constructSelf constructs the environment variable in-place.
func (v *FromEnvironment[V, T]) constructSelf(ctx *Context) error {
	name := identifier.ToScreamingSnakeCase(
		typeOf[V]().Name(),
	)

	raw, ok := os.LookupEnv(name)
	if !ok {
		return fmt.Errorf(
			"the %s environment variable is not defined",
			name,
		)
	}

	if raw == "" {
		return fmt.Errorf(
			"the %s environment variable is defined, but it is empty",
			name,
		)
	}

	if err := parseInto(raw, &v.value); err != nil {
		return fmt.Errorf(
			"the %s environment variable (%#v) is invalid: %w",
			name,
			raw,
			err,
		)
	}

	v.name = name
	v.raw = raw

	return nil
}

// parseInto parses value into *out.
func parseInto(value string, out any) error {
	switch out := out.(type) {
	case *string:
		*out = value

	case *bool:
		return parseBool(value, out)

	case *int:
		return parseInt(value, out, bits.UintSize, math.MinInt, math.MaxInt)
	case *int16:
		return parseInt(value, out, 16, math.MinInt16, math.MaxInt16)
	case *int32:
		return parseInt(value, out, 32, math.MinInt32, math.MaxInt32)
	case *int64:
		return parseInt(value, out, 64, math.MinInt64, math.MaxInt64)

	case *uint:
		return parseUint(value, out, bits.UintSize, math.MaxUint)
	case *uint16:
		return parseUint(value, out, 16, math.MaxUint16)
	case *uint32:
		return parseUint(value, out, 32, math.MaxUint32)
	case *uint64:
		return parseUint(value, out, 64, math.MaxUint64)

	case *float32:
		return parseFloat(value, out, 32)
	case *float64:
		return parseFloat(value, out, 64)

	default:
		panic(fmt.Sprintf(
			"%s implements the Parseable constraint, but is not handled by the parser",
			reflect.TypeOf(out).Elem(),
		))
	}

	return nil
}

// parseBool parses a boolean and assigns it to *out.
func parseBool(value string, out *bool) error {
	switch strings.ToLower(value) {
	case "true", "yes", "on":
		*out = true
	case "false", "no", "off":
		*out = false
	default:
		return errors.New(`expected one of "true", "false", "yes", "no", "on" or "off"`)
	}

	return nil
}

// parseInt parses a signed integer and assigns it to *out.
func parseInt[T constraints.Signed](value string, out *T, size int, min, max T) error {
	n, err := strconv.ParseInt(value, 0, size)
	if err != nil {
		return fmt.Errorf("expected integer between %d and %d", min, max)
	}

	*out = T(n)
	return nil
}

// parseUint parses an unsigned integer and assigns it to *out.
func parseUint[T constraints.Unsigned](value string, out *T, size int, max T) error {
	n, err := strconv.ParseUint(value, 0, size)
	if err != nil {
		return fmt.Errorf("expected integer between 0 and %d", max)
	}

	*out = T(n)
	return nil
}

// parseFloat parses a float and assigns it to *out.
func parseFloat[T constraints.Float](value string, out *T, size int) error {
	n, err := strconv.ParseFloat(value, size)
	if err != nil {
		return errors.New("expected floating-point number")
	}

	*out = T(n)
	return nil
}
