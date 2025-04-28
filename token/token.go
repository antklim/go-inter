package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"

	BANG      TokenType = "!"
	PERIOD    TokenType = "."
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	SLASH     TokenType = "/"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"
	LT     TokenType = "<"
	GT     TokenType = ">"

	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	RETURN   TokenType = "RETURN"

	HUG TokenType = "🤗"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookupIdent(s string) TokenType {
	if tt, ok := keywords[s]; ok {
		return tt
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}
