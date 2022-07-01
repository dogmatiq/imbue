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

			if !more || !strings.HasPrefix(fr.Function, "github.com/dogmatiq/imbue.") {
				return location{fr.File, fr.Line}
			}
		}
	}
}
