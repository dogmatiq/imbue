package location

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Location is a location within the application's source code.
type Location struct {
	File string
	Line int
}

func (l Location) String() string {
	return fmt.Sprintf(
		"%s:%d",
		filepath.Base(l.File),
		l.Line,
	)
}

// Find returns the file and line number of the first frame in the
// current goroutine's stack that is NOT part of the imbue package.
func Find() Location {
	var pointers [8]uintptr
	skip := 2 // Always skip runtime.Callers() and Find().

	for {
		count := runtime.Callers(skip, pointers[:])
		iter := runtime.CallersFrames(pointers[:count])
		skip += count

		for {
			fr, more := iter.Next()

			if !more || !strings.HasPrefix(fr.Function, "github.com/dogmatiq/imbue") {
				return Location{fr.File, fr.Line}
			}
		}
	}
}
