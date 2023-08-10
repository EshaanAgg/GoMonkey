package lexer

import "eshaanagg/go/monkey/token"

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isValidIdentifierCharacter(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_'
}

var keywords = map[string]token.TokenType{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"return": token.RETURN,
	"if":     token.IF,
	"else":   token.ELSE,
}

func getKeyword(key string) token.TokenType {
	tok, found := keywords[key]
	if found {
		return tok
	}
	return token.IDENTIFIER
}
