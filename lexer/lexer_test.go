package lexer

import (
	"errors"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input  string
		tokens []Token
		err    error
	}{
		{
			input: "name EQ 'Bob' AND (age GT 30 OR age LT 20)",
			tokens: []Token{
				{FIELD, "name"},
				{OPERATOR, "EQ"},
				{VALUE, "Bob"},
				{LOGICAL, "AND"},
				{LPAREN, "("},
				{FIELD, "age"},
				{OPERATOR, "GT"},
				{VALUE, "30"},
				{LOGICAL, "OR"},
				{FIELD, "age"},
				{OPERATOR, "LT"},
				{VALUE, "20"},
				{RPAREN, ")"},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		tokens, err := Tokenize(tt.input)
		if !reflect.DeepEqual(tokens, tt.tokens) || !errors.Is(err, tt.err) {
			t.Errorf("Tokenize(%q) = %v, %v; want %v, %v", tt.input, tokens, err, tt.tokens, tt.err)
		}
	}
}
