package lexer

import (
	"github.com/abhinav-0401/hyde/util"
)

type Lexer struct {
	src    string
	start  int
	curr   int
	line   int
	Tokens []Token
}

var keywords = map[string]TokenKind{
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

func New(src string) *Lexer {
	return &Lexer{src: src, line: 1, Tokens: make([]Token, 0)}
}

func (l *Lexer) LexTokens() {
	for !l.isAtEnd() {
		l.start = l.curr
		l.lexToken()
	}

	l.Tokens = append(l.Tokens, Token{Kind: Eof, Lexeme: "EOF", Line: l.line})
}

func (l *Lexer) lexToken() {
	var ch byte = l.advance()

	switch ch {
	case '(':
		l.addToken(LeftParen)
	case ')':
		l.addToken(RightParen)
	case '{':
		l.addToken(LeftBrace)
	case '}':
		l.addToken(RightBrace)
	case ',':
		l.addToken(Comma)
	case '.':
		l.addToken(Dot)
	case '-':
		l.addToken(Minus)
	case '+':
		l.addToken(Plus)
	case ';':
		l.addToken(Semicolon)
	case '*':
		l.addToken(Star)

	case '!':
		if l.match('=') {
			l.addToken(BangEq)
		} else {
			l.addToken(Bang)
		}
	case '=':
		if l.match('=') {
			l.addToken(EqEq)
		} else {
			l.addToken(Eq)
		}
	case '<':
		if l.match('=') {
			l.addToken(LessEq)
		} else {
			l.addToken(Less)
		}
	case '>':
		if l.match('=') {
			l.addToken(GreaterEq)
		} else {
			l.addToken(Greater)
		}
	case '/':
		if l.match('/') {
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
		} else {
			l.addToken(Slash)
		}

	case ' ':
		fallthrough
	case '\r':
		fallthrough
	case '\t': // left empty

	case '\n':
		l.line++

	case '"':
		l.createStrToken()

	default:
		if l.isDigit(ch) {
			l.createNumToken()
		} else if l.isAlpha(ch) {
			l.createIdentToken()
		} else {
			util.Error(l.line, "Unexpected character")
		}
	}
}

func (l *Lexer) createStrToken() {
	var token Token
	token.Kind = String

	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		util.Error(l.line, "Unterminated string\n")
		return
	}

	l.advance()

	token.Lexeme = l.src[l.start:l.curr]
	token.Line = l.line

	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) createNumToken() {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance() // for the . char
		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	var token = Token{Kind: Number, Lexeme: l.src[l.start:l.curr], Line: l.line}
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) createIdentToken() {
	for l.isAlphaNum(l.peek()) && !l.isAtEnd() {
		l.advance()
	}

	var token Token
	token.Lexeme = l.src[l.start:l.curr]
	token.Line = l.line

	kwType, ok := keywords[token.Lexeme]

	if ok { // ident is a keyword
		token.Kind = kwType
	} else {
		token.Kind = Ident
	}

	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) isAlphaNum(ch byte) bool {
	return l.isAlpha(ch) || l.isDigit(ch)
}

func (l *Lexer) addToken(kind TokenKind) {
	var token Token = Token{
		Kind:   kind,
		Lexeme: l.src[l.start:l.curr],
		Line:   l.line,
	}
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) match(ch byte) bool {
	if l.isAtEnd() {
		return false
	}

	if l.src[l.curr] != ch {
		return false
	}

	l.curr++
	return true
}

func (l *Lexer) advance() byte {
	var ch = l.src[l.curr]
	l.curr++
	return ch
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	var ch = l.src[l.curr]
	return ch
}

func (l *Lexer) peekNext() byte {
	if l.curr+1 >= len(l.src) {
		return 0
	}

	var ch = l.src[l.curr+1]
	return ch
}

func (l *Lexer) isAtEnd() bool {
	return l.curr >= len(l.src)
}
