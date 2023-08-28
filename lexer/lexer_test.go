package lexer

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "name EQ 'Bob' AND (age GT 30 OR age LT 20)",
			expected: []Token{
				{Type: FIELD, Value: "name"},
				{Type: OPERATOR, Value: "EQ"},
				{Type: VALUE, Value: "Bob"},
				{Type: LOGICAL, Value: "AND"},
				{Type: LPAREN, Value: "("},
				{Type: FIELD, Value: "age"},
				{Type: OPERATOR, Value: "GT"},
				{Type: VALUE, Value: "30"},
				{Type: LOGICAL, Value: "OR"},
				{Type: FIELD, Value: "age"},
				{Type: OPERATOR, Value: "LT"},
				{Type: VALUE, Value: "20"},
				{Type: RPAREN, Value: ")"},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		tokens, err := l.Tokenize()
		if err != nil {
			t.Fatalf("Failed to tokenize input '%s': %s", tt.input, err.Error())
		}

		if len(tokens) != len(tt.expected) {
			t.Errorf("Token count mismatch for input %q. Got: %d, Expected: %d", tt.input, len(tokens), len(tt.expected))
			continue
		}

		for i := range tokens {
			if tokens[i] != tt.expected[i] {
				t.Errorf("Tokenize(%q) at index %d = %s; want %s", tt.input, i, tokens[i].PrintToken(), tt.expected[i].PrintToken())
			}
		}
	}
}
