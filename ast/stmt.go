package ast

type StmtVisitor interface {
	VisitExpressionStmtStmt(stmt ExpressionStmt) interface{}
	VisitPrintStmtStmt(stmt PrintStmt) interface{}
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
