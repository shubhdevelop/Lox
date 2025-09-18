package printer

import (
	"fmt"
	"github.com/shubhdevelop/Lox/ast"
	"strings"
)

// AstPrinter implements ast.Visitor
type AstPrinter struct{}

// Ensure AstPrinter implements ast.Visitor at compile time
var _ ast.ExprVisitor = (*AstPrinter)(nil)

// Print converts an expression to its string representation
func (p *AstPrinter) Print(expr ast.Expr) string {
	if expr == nil {
		return "nil"
	}
	result := expr.Accept(p)
	if str, ok := result.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", result)
}

// VisitBinaryExpr handles binary expressions
func (p *AstPrinter) VisitBinaryExpr(expr ast.Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// VisitGroupingExpr handles grouping expressions
func (p *AstPrinter) VisitGroupingExpr(expr ast.Grouping) interface{} {
	return p.parenthesize("group", expr.Expression)
}

// VisitLiteralExpr handles literal expressions
func (p *AstPrinter) VisitLiteralExpr(expr ast.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

// VisitUnaryExpr handles unary expressions
func (p *AstPrinter) VisitUnaryExpr(expr ast.Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

// parenthesize wraps expressions in parentheses with an operator/name
func (p *AstPrinter) parenthesize(name string, exprs ...ast.Expr) string {
	var builder strings.Builder

	builder.WriteString("(")
	builder.WriteString(name)

	for _, expr := range exprs {
		if expr == nil {
			builder.WriteString(" nil")
			continue
		}

		builder.WriteString(" ")
		result := expr.Accept(p)
		if str, ok := result.(string); ok {
			builder.WriteString(str)
		} else {
			builder.WriteString(fmt.Sprintf("%v", result))
		}
	}

	builder.WriteString(")")
	return builder.String()
}
