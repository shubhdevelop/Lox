package environment

import (
	"fmt"
	loxErrors "github.com/shubhdevelop/Lox/LoxErrors"
	"github.com/shubhdevelop/Lox/Token"
)

type Environment struct {
	Enclosing *Environment
	Values    map[string]interface{}
}

// Equivalent to Environment() in Java
func NewEnvironment() *Environment {
	return &Environment{
		Enclosing: nil,
		Values:    make(map[string]interface{}),
	}
}

// Equivalent to Environment(Environment enclosing) in Java
func NewEnclosedEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		Enclosing: enclosing,
		Values:    make(map[string]interface{}),
	}
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

func (e *Environment) Get(name token.Token) (interface{}, error) {
	if value, exists := e.Values[name.Lexeme]; exists {
		return value, nil
	}

	if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	}

	error := loxErrors.RuntimeError{
		Token:   name,
		Message: fmt.Sprintf("Undefined variable '%s'.", name.Lexeme),
	}

	panic(error.ThrowRuntimeError())
}

func (e *Environment) Assign(name token.Token, value interface{}) {
	if _, ok := e.Values[name.Lexeme]; ok {
		e.Values[name.Lexeme] = value
		return
	}

	if e.Enclosing != nil {
		e.Enclosing.Assign(name, value)
		return
	}

	error := loxErrors.RuntimeError{
		Token:   name,
		Message: "Undefined variable '" + name.Lexeme + "'.",
	}
	error.ThrowRuntimeError()

}
