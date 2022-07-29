package ferrite

import "unsafe"

// bits returns the number of bits used to store a represent of type T.
func bits[T any]() int {
	var zero T
	return int(unsafe.Sizeof(zero)) * 8
}
