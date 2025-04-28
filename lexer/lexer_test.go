package lexer_test

import (
	"testing"

	"github.com/antklim/go-inter/lexer"
	"github.com/antklim/go-inter/token"
)

func TestNextToken(t *testing.T) {
	t.Run("simple input", func(t *testing.T) {
		input := "=+{}(),🤗;"

		testCases := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.PLUS, expectedLiteral: "+"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.HUG, expectedLiteral: "🤗"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		}

		l := lexer.New(input)

		for i, tC := range testCases {
			tok := l.NextToken()

			if tok.Type != tC.expectedType {
				t.Errorf("test #%d wrong token type: want %q, got %q", i, tC.expectedType, tok.Type)
			}

			if tok.Literal != tC.expectedLiteral {
				t.Errorf("test #%d wrong literal: want %q, got %q", i, tC.expectedLiteral, tok.Literal)
			}
		}
	})

	t.Run("full code input", func(t *testing.T) {
		input := `let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let sub = fn(x, y) {
			return x - y;
		};

		let add_result = add(five, ten);
		let sub_result = sub(7, 1);
		`

		testCases := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "five"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.INT, expectedLiteral: "5"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "ten"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.INT, expectedLiteral: "10"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "add"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.FUNCTION, expectedLiteral: "fn"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.PLUS, expectedLiteral: "+"},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "sub"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.FUNCTION, expectedLiteral: "fn"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.RETURN, expectedLiteral: "return"},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.MINUS, expectedLiteral: "-"},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "add_result"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.IDENT, expectedLiteral: "add"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.IDENT, expectedLiteral: "five"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.IDENT, expectedLiteral: "ten"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "sub_result"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.IDENT, expectedLiteral: "sub"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.INT, expectedLiteral: "7"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.INT, expectedLiteral: "1"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},

			{expectedType: token.EOF, expectedLiteral: string(rune(0))},
		}

		l := lexer.New(input)

		for i, tC := range testCases {
			tok := l.NextToken()

			if tok.Type != tC.expectedType {
				t.Errorf("test #%d wrong token type: want %q, got %q", i, tC.expectedType, tok.Type)
			}

			if tok.Literal != tC.expectedLiteral {
				t.Errorf("test #%d wrong literal: want %q, got %q", i, tC.expectedLiteral, tok.Literal)
			}
		}
	})
}
