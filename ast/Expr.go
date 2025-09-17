package ast

import (
    "github.com/shubhdevelop/Lox/token"
)

type Visitor interface {
    VisitBinaryExpr(expr Binary) interface{}
    VisitGroupingExpr(expr Grouping) interface{}
    VisitLiteralExpr(expr Literal) interface{}
    VisitUnaryExpr(expr Unary) interface{}
}

type Expr interface {
    Accept(visitor Visitor) interface{}
}

type Binary struct {
    Left Expr
    Operator token.Token
    Right Expr
}

func (e Binary) Accept(visitor Visitor) interface{} {
    return visitor.VisitBinaryExpr(e)
}

type Grouping struct {
    Expression Expr
}

func (e Grouping) Accept(visitor Visitor) interface{} {
    return visitor.VisitGroupingExpr(e)
}

type Literal struct {
    Value interface{}
}

func (e Literal) Accept(visitor Visitor) interface{} {
    return visitor.VisitLiteralExpr(e)
}

type Unary struct {
    Operator token.Token
    Right Expr
}

func (e Unary) Accept(visitor Visitor) interface{} {
    return visitor.VisitUnaryExpr(e)
}

