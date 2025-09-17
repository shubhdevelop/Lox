package parser

import (
	"errors"
	"github.com/shubhdevelop/Lox/ast"
	"github.com/shubhdevelop/Lox/loxErrors"
	"github.com/shubhdevelop/Lox/token"
)

type Parser struct {
	current int
	Tokens  []token.Token
}

func (p *Parser) error(token token.Token, message string) error {
	loxErrors.Error(token, message)
	return errors.New("Error while parsing")
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case token.CLASS:
		case token.FUN:
		case token.VAR:
		case token.FOR:
		case token.IF:
		case token.WHILE:
		case token.PRINT:
		case token.RETURN:
			return
		}
		p.advance()
	}
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr

}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tokenType token.TokenType, message string) token.Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(p.error(p.peek(), message))
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.current-1]
}

func (p *Parser) comparison() ast.Expr {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(token.PLUS, token.MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.Binary{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.Unary{
			Operator: operator,
			Right:    right,
		}
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	switch {
	case p.match(token.FALSE):
		return ast.Literal{Value: false}
	case p.match(token.TRUE):
		return ast.Literal{Value: true}
	case p.match(token.NIL):
		return ast.Literal{Value: nil}
	case p.match(token.NUMBER, token.STRING):
		return ast.Literal{Value: p.previous().Lexeme}
	case p.match(token.LEFT_PAREN):
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return ast.Grouping{Expression: expr}
	default:
		panic(p.error(p.peek(), "Expected expression"))
	}
}

func (p *Parser) Parse() ast.Expr {
	expr := p.expression()
	return expr
}
