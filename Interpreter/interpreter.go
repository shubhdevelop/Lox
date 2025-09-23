package interpreter

import (
	"fmt"
	"strconv"

	"github.com/shubhdevelop/Lox/LoxErrors"
	"github.com/shubhdevelop/Lox/ast"
	"github.com/shubhdevelop/Lox/environment"
	"github.com/shubhdevelop/Lox/token"
)

type Interpreter struct{}

var env = environment.Environment{}

var _ ast.ExprVisitor = (*Interpreter)(nil)
var _ ast.StmtVisitor = (*Interpreter)(nil)

func (i *Interpreter) VisitBinaryExpr(expr ast.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL:
		l, lok := left.(float64)
		r, rok := right.(float64)
		if !lok || !rok {
			panic("Operands must be numbers.") // or return an error type
		}

		switch expr.Operator.Type {
		case token.GREATER:
			i.checkNumberOperand(expr.Operator, left, right)
			return l > r
		case token.GREATER_EQUAL:
			i.checkNumberOperand(expr.Operator, left, right)
			return l >= r
		case token.LESS:
			i.checkNumberOperand(expr.Operator, left, right)
			return l < r
		case token.LESS_EQUAL:
			i.checkNumberOperand(expr.Operator, left, right)
			return l <= r
		}
	case token.BANG_EQUAL:
		return !i.isEqual(left, right)
	case token.EQUAL_EQUAL:
		return i.isEqual(left, right)
	case token.MINUS:
		return left.(float64) - right.(float64)
	case token.PLUS:
		runtimeError := loxErrors.RuntimeError{
			Token:   expr.Operator,
			Message: "Operands must be two numbers or two strings.",
		}
		switch l := left.(type) {
		case float64:
			if r, ok := right.(float64); ok {
				return l + r
			} else {
				panic(runtimeError.ThrowRuntimeError())
			}
		case string:
			if r, ok := right.(string); ok {
				return l + r
			} else {
				panic(runtimeError.ThrowRuntimeError())
			}

		}
	case token.SLASH:
		i.checkNumberOperand(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case token.STAR:
		i.checkNumberOperand(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	}
	return nil
}

func (i *Interpreter) VisitGroupingExpr(expr ast.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitLiteralExpr(expr ast.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitUnaryExpr(expr ast.Unary) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case token.BANG:
		return !i.isTruthy(right)
	case token.MINUS:
		return -right.(float64)
	}
	return nil
}

func (i *Interpreter) evaluate(expr ast.Expr) interface{} {
	return expr.Accept(i)
}

func (i *Interpreter) isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}
	if b, ok := object.(bool); ok {
		return b
	}
	return true
}

func (i *Interpreter) isEqual(a, b interface{}) bool {
	// Handle nil explicitly
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Use Go's == only if types are comparable
	switch x := a.(type) {
	case bool:
		if y, ok := b.(bool); ok {
			return x == y
		}
	case float64:
		if y, ok := b.(float64); ok {
			return x == y
		}
	case string:
		if y, ok := b.(string); ok {
			return x == y
		}
	}

	return false // types don't match or not comparable
}

func (i *Interpreter) Interpret(stmts []ast.Stmt) {
	defer func() {
		if r := recover(); r != nil {
			// Equivalent to catching RuntimeError in Java
			fmt.Println("Runtime error:", r)
		}
	}()

	for _, stmt := range stmts {
		i.execute(stmt)
	}
}

func (i *Interpreter) execute(stmt ast.Stmt) {
	stmt.Accept(i)
}

// stringify matches Lox semantics
func stringify(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	switch v := obj.(type) {
	case float64:
		// Convert float to string, but trim trailing ".0"
		text := strconv.FormatFloat(v, 'f', -1, 64)
		return text
	case bool:
		// Lox prints booleans lowercase
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (i Interpreter) checkNumberOperand(operator token.Token, left, right interface{}) {
	runtimeError := loxErrors.RuntimeError{
		Token:   operator,
		Message: "Operands must be numbers.",
	}

	if _, ok := left.(float64); ok {
		if _, ok2 := right.(float64); ok2 {
			return
		}
	}
	panic(runtimeError.ThrowRuntimeError())
}

// Statement Visitors

func (i *Interpreter) VisitExpressionStmtStmt(stmt ast.ExpressionStmt) interface{} {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitPrintStmtStmt(stmt ast.PrintStmt) interface{} {
	value := i.evaluate(stmt.Expression)
	fmt.Println(stringify(value))
	return nil
}

func (i *Interpreter) VisitVarStmtStmt(stmt ast.VarStmt) interface{} {
	var value interface{} = nil

	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}

	env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitVariableExpr(expr ast.Variable) interface{} {
	value, _ := env.Get(expr.Name)
	return value
}
