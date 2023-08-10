package lexer

import (
	"eshaanagg/go/monkey/token"
	"testing"
)

func TestSimpleTokens(t *testing.T) {
	input := `=+-*/%(){},;~`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Operators
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
		{token.MOD, "%"},

		// Delimeters
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},

		// Illegal and EOF character
		{token.ILLEGAL, "~"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}