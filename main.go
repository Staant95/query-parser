package main

import "github.com/Staant95/query-parser/lexer"

func main() {
	// expected from lexer:  FIELD(name) OPERATOR(EQ) STRING("Bob") AND ( FIELD(age) OPERATOR(GT) INTEGER(30) OR FIELD(age) OPERATOR(LT) INTEGER(20) )
	source := "name EQ 'Bob' AND (age GT 30 OR age LT 20)"
	tokens, err := lexer.Tokenize(source)
	if err != nil {
		panic(err)
	}

	for _, t := range tokens {
		println(t.String())
	}
}
