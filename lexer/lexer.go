package lexer

import "eshaanagg/go/monkey/token"

type Lexer struct {
	input        string // The input code that we need to parse
	position     int    // Current position in input (Points to current char)
	readPosition int    // Current reading position in input (After current char)
	ch           byte   // Current char under examination
}

// Factory function to create a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}

	// Initialise the lexer and set read the first character
	l.readChar()
	return l
}

// Read the character at the current position and store the same in the Lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 in ASCII represents NUL which denotes either the EOF or the fact that we haven't read anything
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {

	// Add all the operators
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.MULTIPLY, l.ch)
	case '/':
		tok = token.NewToken(token.DIVIDE, l.ch)
	case '%':
		tok = token.NewToken(token.MOD, l.ch)

	// Delimeters
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)

	// EOF
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		tok = token.NewToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}