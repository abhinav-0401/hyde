package lexer

import (
	"github.com/abhinav-0401/hyde/token"
)

type Lexer struct {
	src     string
	pos     int
	readPos int
	ch      byte
}

func New(src string) *Lexer {
	var l = &Lexer{src: src}
	l.advance()
	return l
}

func (l *Lexer) LexToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = l.newToken(token.LeftParen, l.ch)
	case ')':
		tok = l.newToken(token.RightParen, l.ch)
	case '{':
		tok = l.newToken(token.LeftBrace, l.ch)
	case '}':
		tok = l.newToken(token.RightBrace, l.ch)
	case ',':
		tok = l.newToken(token.Comma, l.ch)
	case '.':
		tok = l.newToken(token.Dot, l.ch)
	case '-':
		tok = l.newToken(token.Minus, l.ch)
	case '+':
		tok = l.newToken(token.Plus, l.ch)
	case ';':
		tok = l.newToken(token.Semicolon, l.ch)
	case '*':
		tok = l.newToken(token.Star, l.ch)
	case '!':
		tok = l.newTwoCharToken('=', token.BangEq, token.Bang)
	case '=':
		tok = l.newTwoCharToken('=', token.EqEq, token.Eq)
	case '>':
		tok = l.newTwoCharToken('=', token.GreaterEq, token.Greater)
	case '<':
		tok = l.newTwoCharToken('=', token.LessEq, token.Less)
	case 0:
		tok.Lexeme = ""
		tok.Kind = token.Eof
	default:
		if l.isLetter(l.ch) {
			tok.Lexeme = l.readIdent()
			tok.Kind = l.lookUpIdent(tok.Lexeme)
			return tok
		} else if l.isDigit(l.ch) {
			tok.Kind = token.Number
			tok.Lexeme = l.readNum()
			return tok
		} else {
			tok = l.newToken(token.Illegal, l.ch)
		}
	}

	l.advance()

	return tok
}

func (l *Lexer) readIdent() string {
	var pos = l.pos
	for l.isAlphaNum() {
		l.advance()
	}

	return l.src[pos:l.pos]
}

func (l *Lexer) readNum() string {
	var pos = l.pos

	for l.isDigit(l.ch) {
		l.advance()
	}

	if l.ch == '.' && l.isDigit(l.peek()) {
		l.advance()
		for l.isDigit(l.ch) {
			l.advance()
		}
	}

	return l.src[pos:l.pos]
}

func (l *Lexer) isAtEnd() bool {
	if l.peek() == 0 {
		return true
	}
	return false
}

func (l *Lexer) advance() {
	if l.readPos >= len(l.src) {
		l.ch = 0
	} else {
		l.ch = l.src[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peek() byte {
	if l.readPos >= len(l.src) {
		return 0
	}
	return l.src[l.readPos]
}

func (l *Lexer) newToken(kind token.TokenKind, ch byte) token.Token {
	return token.Token{Kind: kind, Lexeme: string(ch)}
}

func (l *Lexer) newTwoCharToken(testChar byte, twoCharKind token.TokenKind, oneCharKind token.TokenKind) token.Token {
	var tok token.Token

	if l.peek() == testChar {
		tok.Lexeme = string(l.ch) + string(l.peek())
		tok.Kind = twoCharKind
		l.advance()
	} else {
		tok = l.newToken(oneCharKind, l.ch)
	}

	return tok
}

func (l *Lexer) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) isAlphaNum() bool {
	return (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z') || l.ch == '_' || (l.ch >= '0' && l.ch <= '9')
}

func (l *Lexer) lookUpIdent(lexeme string) token.TokenKind {
	kwType, ok := token.Keywords[lexeme]

	if ok {
		return kwType
	} else {
		return token.Ident
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.advance()
	}
}
