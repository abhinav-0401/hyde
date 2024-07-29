package token

type TokenKind int

type Token struct {
	Kind   TokenKind
	Lexeme string
}

const (
	Eof TokenKind = iota
	Illegal

	// single char tokens
	LeftParen
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
)

var Keywords = map[string]TokenKind{
	"and":    And,
	"struct": Struct,
	"else":   Else,
	"for":    For,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"proc":   Proc,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}
