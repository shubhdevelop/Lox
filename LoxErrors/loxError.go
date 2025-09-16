package loxErrors

import (
	"fmt"
	"os"

	"github.com/shubhdevelop/Lox/state"
)

func report( line int, where , message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

func ThrowNewError(line int, message string) {
	report(line, "", message)
	state.HadError = true
}

