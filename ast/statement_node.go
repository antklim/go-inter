package ast

import (
	"bytes"

	"github.com/antklim/go-inter/token"
)

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode() {}

func (s *LetStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}
	out.WriteRune(';')

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) statementNode() {}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}
	out.WriteRune(';')

	return out.String()
}
