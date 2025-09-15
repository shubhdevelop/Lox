package scanner

import (
	"errors"
	"fmt"
	"strings"
)

type Scanning interface {
	ScanTokens() ([]string, error)
}

type Scanner struct {
	Source string
}


// TokenType represents all possible token types
type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	// End of file
	EOF
)

func (t TokenType) String() string {
	return [...]string{
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE",
		"COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
		"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL",
		"LESS", "LESS_EQUAL",
		"IDENTIFIER", "STRING", "NUMBER",
		"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR",
		"PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE",
		"EOF",
	}[t]
}

type Token struct {
	Type    TokenType   // Enum we defined earlier
	Lexeme  string      // Raw source text
	Literal interface{} // Can hold string, number, etc.
	Line    int         // Line number in source
}

func (t Token) String() string {
	return fmt.Sprintf("%v %s %v", t.Type, t.Lexeme, t.Literal)
}

func (s Scanner) ScanTokens() ([]string, error) {
	if len(s.Source) <= 0 {
		return nil, errors.New("source is empty")
	}
	
	// Initialize empty slice instead of fixed size
	tokens := make([]string, 0)
	words := strings.Fields(s.Source)
	
	for _, word := range words {
		tokens = append(tokens, word)
	}
	
	return tokens, nil
}
