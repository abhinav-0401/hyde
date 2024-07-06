package lexer

import "fmt"

type TokenKind int

type Token struct {
	Kind   TokenKind
	Lexeme string
	Line   int
}

func (t *Token) ToString() {
	fmt.Printf("Kind: %v, Lexeme: %v, Line: %v", t.Kind, t.Lexeme, t.Line)
}

const (
	// single char tokens
	LeftParen TokenKind = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Colon

	// double char tokens
	Bang
	BangEq
	Eq
	EqEq
	Greater
	GreaterEq
	Less
	LessEq

	// literals
	Ident
	String
	Number

	// keywords
	And
	Struct
	Else
	For
	If
	Nil
	Or
	Print
	Proc
	Return
	Super
	This
	True
	Var
	While

	Eof
)

// var kindToString = map[TokenKind]string{
//	LeftParen:  "(",
//	RightParen: ")",
//	LeftBrace:  "{",
//	RightBrace: "}",
//	Comma:      ",",
//	Dot:        ".",
//	Minus:      "-",
//	Plus:       "",
//	Semicolon:  ";",
//	Slash:      "/",
//	Star:       "*",
// }
