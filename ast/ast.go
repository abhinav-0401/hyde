package ast

import (
	"github.com/abhinav-0401/hyde/lexer"
)

type NodeKind string

const (
	BinaryExprNode     NodeKind = "BinaryExprNode"
	UnaryExprNode      NodeKind = "UnaryExprNode"
	GroupingExprNode   NodeKind = "GroupingExprNode"
	NumLiteralExprNode NodeKind = "NumLiteralExprNode"
	StrLiteralExprNode NodeKind = "StrLiteralExprNode"
)

type Expr interface {
	ExprKind() NodeKind

	AcceptPrinter(printer PrinterVisitor) string
	AcceptInterpreter(interpreter InterpreterVisitor)
}

type BinaryExpr struct {
	Left  Expr
	Right Expr
	Op    lexer.Token
}

func (be *BinaryExpr) ExprKind() NodeKind {
	return BinaryExprNode
}

func (be *BinaryExpr) AcceptPrinter(printer PrinterVisitor) string {
	return printer.visitBinaryExpr(be)
}

func (be *BinaryExpr) AcceptInterpreter(interpreter InterpreterVisitor) {
	interpreter.visitBinaryExpr(be)
}

type UnaryExpr struct {
	Op    lexer.Token
	Right Expr
}

func (be *UnaryExpr) ExprKind() NodeKind {
	return UnaryExprNode
}

func (ue *UnaryExpr) AcceptPrinter(printer PrinterVisitor) string {
	return printer.visitUnaryExpr(ue)
}

func (ue *UnaryExpr) AcceptInterpreter(interpreter InterpreterVisitor) {
	interpreter.visitUnaryExpr(ue)
}

type GroupingExpr struct {
	Expression Expr
}

func (ge *GroupingExpr) AcceptPrinter(printer PrinterVisitor) string {
	return printer.visitGroupingExpr(ge)
}

func (ge *GroupingExpr) AcceptInterpreter(interpreter InterpreterVisitor) {
	interpreter.visitGroupingExpr(ge)
}

func (ge *GroupingExpr) ExprKind() NodeKind {
	return GroupingExprNode
}

type NumLiteralExpr struct {
	Value float64
}

type StrLiteralExpr struct {
	Value string // not so sure about this being an interface{} type, but literals can hold a lot of values and Go doesn't give me sum types
}

func (nle *NumLiteralExpr) AcceptPrinter(printer PrinterVisitor) string {
	return printer.visitNumLiteralExpr(nle)
}

func (nle *NumLiteralExpr) AcceptInterpreter(interpreter InterpreterVisitor) {
	interpreter.visitNumLiteralExpr(nle)
}

func (nle *NumLiteralExpr) ExprKind() NodeKind {
	return NumLiteralExprNode
}

func (sle *StrLiteralExpr) AcceptPrinter(printer PrinterVisitor) string {
	return printer.visitStrLiteralExpr(sle)
}

func (sle *StrLiteralExpr) AcceptInterpreter(interpreter InterpreterVisitor) {
	interpreter.visitStrLiteralExpr(sle)
}

func (sle *StrLiteralExpr) ExprKind() NodeKind {
	return StrLiteralExprNode
}
