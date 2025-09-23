package ast

import (
	"github.com/shubhdevelop/Lox/Token"
)

type ExprVisitor interface {
	VisitBinaryExpr(expr Binary) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitUnaryExpr(expr Unary) interface{}
	VisitVariableExpr(expr Variable) interface{}
}

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (n Binary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitBinaryExpr(n)
}

type Grouping struct {
	Expression Expr
}

func (n Grouping) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitGroupingExpr(n)
}

type Literal struct {
	Value interface{}
}

func (n Literal) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitLiteralExpr(n)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (n Unary) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitUnaryExpr(n)
}

type Variable struct {
	Name token.Token
}

func (n Variable) Accept(visitor ExprVisitor) interface{} {
	return visitor.VisitVariableExpr(n)
}
