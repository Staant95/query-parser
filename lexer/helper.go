package lexer

import (
	"bufio"
	"strings"
	"unicode"
)

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		reader: bufio.NewReader(strings.NewReader(input)),
	}
	l.readChar() // initialize with first char, it also sets l.ch
	return l
}
