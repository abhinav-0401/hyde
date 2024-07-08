package ast

import "fmt"

type PrinterVisitor interface {
	visitBinaryExpr(expr *BinaryExpr) string
	visitUnaryExpr(expr *UnaryExpr) string
	visitGroupingExpr(expr *GroupingExpr) string
	visitNumLiteralExpr(expr *NumLiteralExpr) string
	visitStrLiteralExpr(expr *StrLiteralExpr) string
}

type InterpreterVisitor interface {
	visitBinaryExpr(expr *BinaryExpr)
	visitUnaryExpr(expr *UnaryExpr)
	visitGroupingExpr(expr *GroupingExpr)
	visitNumLiteralExpr(expr *NumLiteralExpr)
	visitStrLiteralExpr(expr *StrLiteralExpr)
}

type AstPrinter struct {
}

func (ap *AstPrinter) Print(expr Expr) string {
	return expr.AcceptPrinter(ap)
}

func (ap *AstPrinter) visitBinaryExpr(expr *BinaryExpr) string {
	return ap.parenthesise(expr.Op.Lexeme, expr.Left, expr.Right)
}

func (ap *AstPrinter) visitUnaryExpr(expr *UnaryExpr) string {
	return ap.parenthesise(expr.Op.Lexeme, expr.Right)
}

func (ap *AstPrinter) visitGroupingExpr(expr *GroupingExpr) string {
	return ap.parenthesise("group", expr.Expression)
}

func (ap *AstPrinter) visitNumLiteralExpr(expr *NumLiteralExpr) string {
	return fmt.Sprintf("%v", expr.Value)
}

func (ap *AstPrinter) visitStrLiteralExpr(expr *StrLiteralExpr) string {
	return expr.Value
}

func (ap *AstPrinter) parenthesise(name string, exprs ...Expr) string {
	var pretty string

	pretty += "("
	pretty += name

	for _, expr := range exprs {
		pretty += " "
		pretty += expr.AcceptPrinter(ap)
	}
	pretty += ")"

	return pretty
}
