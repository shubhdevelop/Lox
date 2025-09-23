package ast

import (
	"github.com/shubhdevelop/Lox/Token"
)

type StmtVisitor interface {
	VisitExpressionStmtStmt(stmt ExpressionStmt) interface{}
	VisitPrintStmtStmt(stmt PrintStmt) interface{}
	VisitVarStmtStmt(stmt VarStmt) interface{}
}

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitExpressionStmtStmt(n)
}

type PrintStmt struct {
	Expression Expr
}

func (n PrintStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitPrintStmtStmt(n)
}

type VarStmt struct {
	Name        token.Token
	Initializer Expr
}

func (n VarStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.VisitVarStmtStmt(n)
}
