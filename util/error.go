package util

import (
	"fmt"
	"os"
)

func Error(line int, msg string) {
	Report(line, "", msg)
}

func Report(line int, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %v ] Error%v: %v", line, where, msg)
	StateFlags.HadErr = true
}
