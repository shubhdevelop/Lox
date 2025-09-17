package loxErrors

import (
	"fmt"
	"os"

	"github.com/shubhdevelop/Lox/state"
	"github.com/shubhdevelop/Lox/token"
)

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

func Error(token_p token.Token, message string) {
	// token_p because the it's clashing the token module name
	// _p suggest the the parameter
	if token_p.Type == token.EOF {
		report(token_p.Line, " at end", message)
	} else {
		report(token_p.Line, " at '"+token_p.Lexeme+"'", message)
	}

}

func ThrowNewError(line int, message string) {
	report(line, "", message)
	state.HadError = true
}
