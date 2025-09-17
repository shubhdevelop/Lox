package parser

import (
	"github.com/shubhdevelop/Lox/token"
	"github.com/shubhdevelop/Lox/ast"
)

type Parser struct {
	current int
	Tokens []token.Token
}

func (p *Parser)  expression() ast.Expr {
	return p.equality()
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL){
		operator := p.previous()
		right := p.comparison()
		expr = ast.Binary{
			Left:expr,
			Operator: operator,
			Right: right
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

func (p *Parser) check(type token.TokenType) bool {
	if p.isAtEnd(){
		return false
	}
	return p.peek().type == type
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd(){
		p.current++;
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.current] 
}

func (p *Parser) previous() {
	return p.Tokens[p.current - 1];
}


func (p *Parser) comparison() {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.expression()
		right = p.factor()
		expr = ast.Binary{Left:expr, Operator:  operator, Right:  right} 
	}
	return expr 
}

func (p *Parser) term() {
	expr := p.factor()
	for p.match(token.PLUS, token.MINUS) {
		operator := p.expression()
		right = p.factor()
		expr = ast.Binary{Left:expr, Operator:  operator, Right:  right} 
	}
	return expr 
}


func (p *Parser) factor() {
	expr := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.expression()
		right = p.unary()
		expr = ast.Binary{Left:expr, Operator:  operator, Right:  right} 
	}
	return expr 
}

func (p *Parser) unary() {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.urinary()
		return ast.Unary{
			Operator: operator,
			Right: right,
		}
	}
	return p.primary()
}

func (p *Parser) primary(){
	switch {
	case p.match(token.FALSE):
		return ast.Literal{false}
	case p.match(token.TRUE):
		return ast.Literal{true}
	case p.match(token.NIL):
		return ast.Literal{nil}
	case p.match(token.NUMBER, token.STRING):
		return ast.Literal{p.previous().literal}
	case p.match(token.LEFT_PAREN):
		expr := p.expression()
		p.consume(token.RIGHT_PAREN,"Expect ')' after expression.") 
		return ast.Grouping{expr}
	default:
	}
}

