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
		l.ch = 0 // 0 in ASCII represents NULL which denotes either the EOF or the fact that we haven't read anything
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isValidIdentifierCharacter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {

	// Operators
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQUALS, Literal: "=="}
			l.readChar()
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '%':
		tok = token.NewToken(token.MOD, l.ch)

	case '>':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.GTE, Literal: ">="}
			l.readChar()
		} else {
			tok = token.NewToken(token.GT, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.LTE, Literal: "<="}
			l.readChar()
		} else {
			tok = token.NewToken(token.LT, l.ch)
		}

	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOTEQUALS, Literal: "!="}
			l.readChar()
		} else {
			tok = token.NewToken(token.NOT, l.ch)
		}

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

	// End of File
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(l.ch) {
			identifier := l.readIdentifier()
			tok = token.Token{Type: getKeyword(identifier), Literal: identifier}
			return tok
		} else if isDigit(l.ch) {
			tok = token.Token{Type: token.INT, Literal: l.readNumber()}
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
