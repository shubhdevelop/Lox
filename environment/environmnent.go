package environment

import (
	"fmt"
	loxErrors "github.com/shubhdevelop/Lox/LoxErrors"
	"github.com/shubhdevelop/Lox/Token"
)

type Environment struct {
	Values map[string]interface{}
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

func (e *Environment) Get(name token.Token) (interface{}, error) {
	if value, exists := e.Values[name.Lexeme]; exists {
		return value, nil
	}

	error := loxErrors.RuntimeError{
		Token:   name,
		Message: fmt.Sprintf("Undefined variable '%s'.", name.Lexeme),
	}

	panic(error.ThrowRuntimeError())
}
