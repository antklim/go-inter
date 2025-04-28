package lexer

import (
	"unicode/utf8"

	"github.com/antklim/go-inter/token"
)

type Lexer struct {
	input          string
	ch             rune
	chPosition     int
	nextChPosition int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newRuneToken(token.ASSIGN, l.ch)
	case '+':
		tok = newRuneToken(token.PLUS, l.ch)
	case '-':
		tok = newRuneToken(token.MINUS, l.ch)
	case ',':
		tok = newRuneToken(token.COMMA, l.ch)
	case ';':
		tok = newRuneToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newRuneToken(token.LPAREN, l.ch)
	case ')':
		tok = newRuneToken(token.RPAREN, l.ch)
	case '{':
		tok = newRuneToken(token.LBRACE, l.ch)
	case '}':
		tok = newRuneToken(token.RBRACE, l.ch)
	case '🤗':
		tok = newRuneToken(token.HUG, l.ch)
	case 0:
		tok = newRuneToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			l := l.readIdentifier()
			t := token.LookupIdent(l)
			return newToken(t, l)
		} else if isDigit(l.ch) {
			l := l.readNumber()
			tok = newToken(token.INT, l)
			return tok
		} else {
			tok = newRuneToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextChPosition >= len(l.input) {
		l.ch = 0
		return
	}

	ch, w := utf8.DecodeRuneInString(l.input[l.nextChPosition:])
	l.ch = ch
	l.chPosition = l.nextChPosition
	l.nextChPosition += w
}

func (l *Lexer) readIdentifier() string {
	return l.readString(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readString(isDigit)
}

// readString reads and returns string as long predicate function f returns true.
func (l *Lexer) readString(f func(rune) bool) string {
	position := l.chPosition
	for f(l.ch) {
		l.readChar()
	}
	return l.input[position:l.chPosition]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'

}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func newToken(tt token.TokenType, l string) token.Token {
	return token.Token{Type: tt, Literal: l}
}

func newRuneToken(tt token.TokenType, ch rune) token.Token {
	return newToken(tt, string(ch))
}
