package lexer

import (
	"bufio"
	"unicode"
)

const (
	FIELD TokenType = iota
	OPERATOR
	VALUE
	LOGICAL
	LPAREN
	RPAREN
)

type Lexer struct {
	input  string
	pos    int  // current position in input
	ch     rune // current char
	reader *bufio.Reader
}

// Define all keywords
var logicalOperators = []string{"AND", "OR"}
var comparisonOperators = []string{"EQ", "GT", "LT"}

// Tokenize Add Tokenize function to Lexer struct
func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token

	for l.ch != 0 {
		switch {
		// skip whitespace
		case unicode.IsSpace(l.ch):
			l.readChar()
		case l.ch == '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
			l.readChar()
		case l.ch == ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
			l.readChar()
		case l.ch == '\'':
			tokens = append(tokens, l.readStringValue())
		default:
			tokens = append(tokens, l.readWord())
		}
	}

	return tokens, nil
}

func (l *Lexer) classifyWord(word string) Token {
	if contains(logicalOperators, word) {
		return Token{Type: LOGICAL, Value: word}
	} else if contains(comparisonOperators, word) {
		return Token{Type: OPERATOR, Value: word}
	} else if isNumeric(word) {
		return Token{Type: VALUE, Value: word}
	} else {
		return Token{Type: FIELD, Value: word}
	}
}

func (l *Lexer) readChar() {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		l.ch = 0 // EOF
	} else {
		l.ch = ch
	}
	l.pos++
}

func (l *Lexer) readStringValue() Token {
	l.readChar() // skip opening quote
	startPos := l.pos

	for l.ch != '\'' && l.ch != 0 {
		l.readChar()
	}

	// l.ch == 0 is EOF
	if l.ch == 0 {
		panic("unterminated string value")
	}

	value := l.input[startPos-1 : l.pos-1]
	l.readChar() // skip closing quote
	return Token{Type: VALUE, Value: value}
}

func (l *Lexer) readWord() Token {
	startPos := l.pos

	// l.ch == 0 is EOF
	// a word is a sequence of non-space characters and cannot contain the following characters:
	// (
	// )
	// '
	for !unicode.IsSpace(l.ch) && l.ch != '(' && l.ch != ')' && l.ch != 0 && l.ch != '\'' {
		l.readChar()
	}

	word := l.input[startPos-1 : l.pos-1]
	return l.classifyWord(word)
}
