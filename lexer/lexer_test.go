package lexer_test

import (
	"testing"

	"github.com/antklim/go-inter/lexer"
	"github.com/antklim/go-inter/token"
)

func TestNextToken(t *testing.T) {
	t.Run("single characters tokenisation", func(t *testing.T) {
		input := "=+{}(),🤗;!-/*5 < 10 > 8.,"

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
			{expectedType: token.BANG, expectedLiteral: "!"},
			{expectedType: token.MINUS, expectedLiteral: "-"},
			{expectedType: token.SLASH, expectedLiteral: "/"},
			{expectedType: token.ASTERISK, expectedLiteral: "*"},
			{expectedType: token.INT, expectedLiteral: "5"},
			{expectedType: token.LT, expectedLiteral: "<"},
			{expectedType: token.INT, expectedLiteral: "10"},
			{expectedType: token.GT, expectedLiteral: ">"},
			{expectedType: token.INT, expectedLiteral: "8"},
			{expectedType: token.PERIOD, expectedLiteral: "."},
			{expectedType: token.COMMA, expectedLiteral: ","},
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

	t.Run("assignments and functions", func(t *testing.T) {
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

	t.Run("if statements", func(t *testing.T) {
		input := `
		let gt = fn(x, y) {
			if x > y {
				return true;
			} else {
				return false;
			}
		}

		let negative = fn(x) {
			x < 0
		}
		`

		testCases := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "gt"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.FUNCTION, expectedLiteral: "fn"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.COMMA, expectedLiteral: ","},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.IF, expectedLiteral: "if"},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.GT, expectedLiteral: ">"},
			{expectedType: token.IDENT, expectedLiteral: "y"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.RETURN, expectedLiteral: "return"},
			{expectedType: token.TRUE, expectedLiteral: "true"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},
			{expectedType: token.ELSE, expectedLiteral: "else"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.RETURN, expectedLiteral: "return"},
			{expectedType: token.FALSE, expectedLiteral: "false"},
			{expectedType: token.SEMICOLON, expectedLiteral: ";"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},

			{expectedType: token.LET, expectedLiteral: "let"},
			{expectedType: token.IDENT, expectedLiteral: "negative"},
			{expectedType: token.ASSIGN, expectedLiteral: "="},
			{expectedType: token.FUNCTION, expectedLiteral: "fn"},
			{expectedType: token.LPAREN, expectedLiteral: "("},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.RPAREN, expectedLiteral: ")"},
			{expectedType: token.LBRACE, expectedLiteral: "{"},
			{expectedType: token.IDENT, expectedLiteral: "x"},
			{expectedType: token.LT, expectedLiteral: "<"},
			{expectedType: token.INT, expectedLiteral: "0"},
			{expectedType: token.RBRACE, expectedLiteral: "}"},

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
