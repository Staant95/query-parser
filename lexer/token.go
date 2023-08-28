package lexer

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) PrintToken() string {
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
