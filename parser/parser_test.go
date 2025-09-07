package parser_test

import (
	"testing"

	"github.com/antklim/go-inter/ast"
	"github.com/antklim/go-inter/lexer"
	"github.com/antklim/go-inter/parser"
)

func TestParseLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	testCases := []struct {
		want string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tc := range testCases {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tc.want)
	}
}

func TestParseReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 838383;
`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, got %d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Errorf("stmt.Expression is not *ast.Identifier, got %T", stmt.Expression)
	}

	if want, got := "foobar", ident.Value; want != got {
		t.Errorf("invalid ident.Value\n\twant %s\n\t got %s", want, got)
	}
	if want, got := "foobar", ident.TokenLiteral(); want != got {
		t.Errorf("invalid ident.TokenLiteral\n\twant %s\n\t got %s", want, got)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, got %d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("stmt.Expression is not *ast.IntegerLiteral, got %T", stmt.Expression)
	}
	if want, got := int64(5), literal.Value; want != got {
		t.Errorf("invalid literal.Value\n\twant %d\n\t got %d", want, got)
	}
	if want, got := "5", literal.TokenLiteral(); want != got {
		t.Errorf("invalid ident.TokenLiteral\n\twant %s\n\t got %s", want, got)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	t.Helper()

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s", name, letStmt.Name.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	t.Helper()

	errs := p.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errs))
	for _, msg := range errs {
		t.Logf("parser error: %q", msg)
	}
	t.FailNow()
}
