package ast_test

import (
	"testing"

	"github.com/antklim/go-inter/ast"
	"github.com/antklim/go-inter/token"
)

func TestString(t *testing.T) {
	p := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if want, got := "let myVar = anotherVar;", p.String(); want != got {
		t.Errorf("invalid program.String()\n\twant: %s\n\t got: %s", want, got)
	}
}
