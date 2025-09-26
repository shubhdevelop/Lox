package scanner

import (
	"errors"
	"github.com/shubhdevelop/Lox/Token"
	"github.com/shubhdevelop/Lox/yaplErrors"
	"strconv"
)

type Scanning interface {
	ScanTokens() ([]string, error)
}

type Scanner struct {
	Source  string
	Tokens  []token.Token
	start   int
	current int
	line    int
}

var KeywordMap = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
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

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if runes := []rune(s.Source); runes[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	runes := []rune(s.Source)
	return runes[s.current]
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.Source) {
		return '\000'
	}
	runes := []rune(s.Source)
	return runes[s.current+1]
}

func (s *Scanner) isAlpha(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_') {
		return true
	}
	return false
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isAlpha(c)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		yaplErrors.ThrowNewError(s.line, "Unterminated string.")
		return
	}
	s.advance()
	value := s.Source[s.start+1 : s.current-1]
	s.addToken(token.STRING, value)
}

func (s *Scanner) isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	value := s.Source[s.start:s.current]
	valueInFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		yaplErrors.ThrowNewError(s.line, "Unexpected Numerical Value")
	}
	s.addToken(token.NUMBER, valueInFloat)
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := string([]rune(s.Source)[s.start:s.current])
	tokenType, ok := KeywordMap[text]
	if !ok {
		s.addToken(token.IDENTIFIER, text)
	} else {
		s.addToken(tokenType, nil)
	}
}

func (s *Scanner) scanToken() {
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
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL, nil)
		} else {
			s.addToken(token.BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL, nil)
		} else {
			s.addToken(token.EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL, nil)
		} else {
			s.addToken(token.LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL, nil)
		} else {
			s.addToken(token.GREATER, nil)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			yaplErrors.ThrowNewError(s.line, "Unexpected Character Encountered")
		}
	}
}

func (s *Scanner) ScanTokens() ([]token.Token, error) {
	if len(s.Source) == 0 {
		return nil, errors.New("source is empty")
	}
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()

	}
	s.Tokens = append(s.Tokens, token.Token{token.EOF, "", nil, s.line})
	return s.Tokens, nil
}
