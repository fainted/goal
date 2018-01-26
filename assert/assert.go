// Package assert provides missing assertion feature in go.
//
// mainly for help writing test code.
package assert

import (
	"log"
	"runtime"
	"strings"
)

var logFunc = log.Fatalf

func assert(v bool, callDepth int, msg string) {
	if v {
		return
	}

	_, file, line, ok := runtime.Caller(callDepth + 1)
	if ok {
		slash := strings.LastIndex(file, "/")
		if slash > 0 {
			file = file[slash+1:]
		}
	}

	logFunc("[%s:%d] %s", file, line, msg)
}

// True do fatal log if v not equals `true`
func True(v bool) {
	assert(v, 1, "assert true")
}

// False do fatal log if v not equals `false`
func False(v bool) {
	assert(!v, 1, "assert false")
}
