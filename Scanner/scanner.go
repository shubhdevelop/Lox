package scanner

import (
	"errors"
	"github.com/shubhdevelop/Lox/token"
	"github.com/shubhdevelop/Lox/loxErrors"
)

type Scanning interface {
	ScanToken() ([]string, error)
}

type Scanner struct {
	Source string
	Tokens []token.Token
	start int
	current int
	line int
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) advance() rune {
	runes := []rune(s.Source)
	ch := runes[s.current]
	s.current++
	return ch
}

func (s *Scanner) addToken(t token.TokenType, literal interface{}) {
	text := s.Source[s.start:s.current] // substring
	tok := token.Token{
		Type:    t,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	}
	s.Tokens = append(s.Tokens, tok)
}

func (s *Scanner) addTokenNoLiteral(t token.TokenType) {
	s.addToken(t, nil)
}


func (s *Scanner) ScanToken() ([]token.Token, error) {
	if len(s.Source) == 0 {
		return nil, errors.New("source is empty")
	}

	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	default: loxErrors.ThrowNewError(s.line, "hello" )
	}
	return s.Tokens, nil
}
