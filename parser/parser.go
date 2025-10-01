package parser

import (
	"errors"

	"github.com/shubhdevelop/YAPL/Token"
	"github.com/shubhdevelop/YAPL/YaplErrors"
	"github.com/shubhdevelop/YAPL/ast"
)

type Parser struct {
	current int
	Tokens  []token.Token
}

func (p *Parser) error(token token.Token, message string) error {
	yaplErrors.Error(token, message)
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
	return p.assignment()
}

func (p *Parser) assignment() ast.Expr {
	expr := p.or()
	if p.match(token.EQUAL) {
		equals := p.previous()
		value := p.assignment()
		if variable, ok := expr.(ast.Variable); ok {
			name := variable.Name
			return ast.Assign{
				Name:  name,
				Value: value,
			}
		}
		yaplErrors.Error(equals, "Invalid assignment target.")
	}
	return expr
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

func (p *Parser) or() ast.Expr {
	expr := p.and()
	for p.match(token.OR) {
		operator := p.previous()
		right := p.and()
		expr = ast.Logical{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) and() ast.Expr {
	expr := p.equality()

	for p.match(token.AND) {
		operator := p.previous()
		right := p.equality()
		expr = ast.Logical{
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
	case p.match(token.NUMBER):
		return ast.Literal{Value: p.previous().Literal}
	case p.match(token.STRING):
		return ast.Literal{Value: p.previous().Literal}
	case p.match(token.IDENTIFIER):
		return ast.Variable{
			Name: p.previous(),
		}

	case p.match(token.LEFT_PAREN):
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return ast.Grouping{Expression: expr}
	default:
		panic(p.error(p.peek(), "Expected expression"))
	}
}

func (p *Parser) Parse() []ast.Stmt {
	statements := []ast.Stmt{}
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	return statements
}

func (p *Parser) declaration() ast.Stmt {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(yaplErrors.RuntimeError); ok {
				p.synchronize()
			} else {
				panic(r) // rethrow if it's not a ParseError
			}
		}
	}()

	if p.match(token.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) varDeclaration() ast.Stmt {
	name := p.consume(token.IDENTIFIER, "Expected variable name")
	var initializer ast.Expr = nil
	if p.match(token.EQUAL) {
		initializer = p.expression()
	}
	p.consume(token.SEMICOLON, "Expect ';' after variable declaration.")
	return ast.VarStmt{Name: name, Initializer: initializer}
}

func (p *Parser) statement() ast.Stmt {
	if p.match(token.FOR) {
		return p.forStatement()
	}
	if p.match(token.IF) {
		return p.ifStatement()
	}
	if p.match(token.PRINT) {
		return p.printStatement()
	}

	if p.match(token.WHILE) {
		return p.whileStatement()
	}
	if p.match(token.LEFT_BRACE) {
		return ast.BlockStmt{
			Statement: p.block(),
		}
	}
	return p.expressionStatement()
}

func (p *Parser) forStatement() ast.Stmt {
	p.consume(token.LEFT_PAREN, "Expect '(' after 'for'.")
	var initializer ast.Stmt
	if p.match(token.SEMICOLON) {
		initializer = nil
	} else if p.match(token.VAR) {
		initializer = p.varDeclaration()
	} else {
		initializer = p.expressionStatement()
	}

	var condition ast.Expr
	if !p.check(token.SEMICOLON) {
		condition = p.expression()
	}
	p.consume(token.SEMICOLON, "Expect ';' after loop condition.")

	var increment ast.Expr
	if !p.check(token.RIGHT_PAREN) {
		increment = p.expression()
	}
	p.consume(token.RIGHT_PAREN, "Expect ')' after for clauses.")

	body := p.statement()

	if increment != nil {
		body = ast.BlockStmt{
			Statement: []ast.Stmt{
				body,
				ast.ExpressionStmt{
					Expression: increment,
				},
			},
		}
	}

	if condition == nil {
		condition = ast.Literal{Value: true}
	}
	body = ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}

	if initializer != nil {
		body = ast.BlockStmt{
			Statement: []ast.Stmt{
				initializer,
				body,
			},
		}
	}

	return body
}

func (p *Parser) ifStatement() ast.Stmt {
	p.consume(token.LEFT_PAREN, "Expect '(' after 'if'.")
	condition := p.expression()

	p.consume(token.RIGHT_PAREN, "Expect ')' after 'if'.")
	thenBranch := p.statement()
	var elseBranch ast.Stmt
	if p.match(token.ELSE) {
		elseBranch = p.statement()
	}
	return ast.IfStmt{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (p *Parser) whileStatement() ast.Stmt {
	p.consume(token.LEFT_PAREN, "Expect '(' after 'while'.")
	condition := p.expression()
	p.consume(token.RIGHT_PAREN, "Expect ')' after condition.")
	body := p.statement()
	return ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}

}

func (p *Parser) block() []ast.Stmt {
	statements := []ast.Stmt{}
	for !p.check(token.RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	p.consume(token.RIGHT_BRACE, "Expect '}' after block.")
	return statements
}

func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return ast.PrintStmt{
		Expression: value,
	}
}

func (p *Parser) expressionStatement() ast.Stmt {
	value := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return ast.ExpressionStmt{
		Expression: value,
	}
}
