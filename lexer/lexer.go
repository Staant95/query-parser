package lexer

import (
	"bufio"
	"errors"
	"strings"
	"unicode"
)

// Define all operators
var logicalOperators = []string{"AND", "OR"}
var comparisonOperators = []string{"EQ", "GT", "LT"}

func classifyWord(word string) TokenType {
	if contains(logicalOperators, word) {
		return LOGICAL
	}
	if contains(comparisonOperators, word) {
		return OPERATOR
	}
	return FIELD
}

type TokenType int

const (
	FIELD TokenType = iota
	OPERATOR
	VALUE
	LOGICAL
	LPAREN
	RPAREN
)

type Token struct {
	Type  TokenType
	Value string
}

// for debugging
// attach function to Token struct
func (t Token) String() string {
	switch t.Type {
	case FIELD:
		return "FIELD:" + t.Value
	case OPERATOR:
		return "OPERATOR:" + t.Value
	case VALUE:
		return "VALUE:" + t.Value
	case LOGICAL:
		return "LOGICAL:" + t.Value
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	default:
		return "UNKNOWN"
	}
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token

	reader := bufio.NewReader(strings.NewReader(input))
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			// End of string, break the loop
			break
		}

		// Skip whitespaces
		if unicode.IsSpace(ch) {
			continue
		}

		if ch == '(' {
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
			continue
		} else if ch == ')' {
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
			continue
		}
		// Handle individual characters and words
		switch ch {
		case '\'': // Starting of a string value
			// Read the string value up to the closing quote
			s, err := reader.ReadString('\'')
			if err != nil {
				// If a closing quote isn't found, it's an error
				return nil, errors.New("unterminated string value")
			}
			// Store the string value without quotes
			tokens = append(tokens, Token{Type: VALUE, Value: s[:len(s)-1]})
		default:
			// Accumulate word characters for operators, logical operators, and fields
			word := string(ch)

			// Example: age GT 30
			// Step 1: word = "a"
			// Step 2: word = "ag"
			// Step 3: word = "age" (breaks loop because nextCh = " ")

			// This for loop is effectively a while(true) loop
			for {
				nextCh, _, err := reader.ReadRune()
				if err != nil {
					// If end of string, break the loop
					break
				}

				if unicode.IsSpace(nextCh) || nextCh == '(' || nextCh == ')' {
					err := reader.UnreadRune()
					if err != nil {
						return nil, err
					}
					break
				}
				word += string(nextCh)
			}

			// Classify the accumulated word token
			if contains(logicalOperators, word) {
				tokens = append(tokens, Token{Type: LOGICAL, Value: word})
			} else if contains(comparisonOperators, word) {
				tokens = append(tokens, Token{Type: OPERATOR, Value: word})
			} else if isNumeric(word) { // Add this check
				tokens = append(tokens, Token{Type: VALUE, Value: word})
			} else {
				tokens = append(tokens, Token{Type: FIELD, Value: word})
			}

		}
	}

	return tokens, nil
}
