package imbue

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// location represents a location within source code.
type location struct {
	File string
	Line int
}

func (l location) String() string {
	return fmt.Sprintf(
		"%s:%d",
		filepath.Base(l.File),
		l.Line,
	)
}

// findLocation returns the file and line number of the first frame in the
// current goroutine's stack that is NOT part of the imbue package.
func findLocation() location {
	var pointers [8]uintptr
	skip := 2 // Always skip runtime.Callers() and findLocation().

	for {
		count := runtime.Callers(skip, pointers[:])
		iter := runtime.CallersFrames(pointers[:count])
		skip += count

		for {
			fr, more := iter.Next()

			if !more || !isImbueFrame(fr) {
				return location{
					fr.File,
					fr.Line,
				}
			}
		}
	}
}

// isImbueFrame returns true if the given frame is part of the imbue package.
//
// Note that we cannot simply use the package path that is prefixed to the
// function name, because as of Go v1.21 it is possible to have a single frame
// referring to both a closure from outside imbue and an (inlined?) generic
// function from within. I'm not sure quite why this is the case, but as a
// result we now use the file path to determine what is and isn't part of the
// imbue package.
func isImbueFrame(fr runtime.Frame) bool {
	return imbueDir != "" &&
		filepath.Dir(fr.File) == imbueDir &&
		!strings.HasSuffix(fr.File, "_test.go")
}

// imbueDir is the absolute path to the imbue package directory.
var imbueDir string

func init() {
	var pointers [1]uintptr

	// Find the stack frame for this init() function by skipping
	// runtime.Callers().
	count := runtime.Callers(1, pointers[:])

	iter := runtime.CallersFrames(pointers[:count])
	fr, _ := iter.Next()
	imbueDir = filepath.Dir(fr.File)
}
