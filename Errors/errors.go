package error
import (
	"github.com/shubhdevelop/Lox/state"
	"fmt"
	"os"
)


func report( line int, where , message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

func ThrowNewError(line int, message string) {
	report(line, "", message)
	state.HadError = true
}
