package main

import "github.com/Staant95/query-parser/lexer"

func main() {
	source := "name EQ 'Bob' AND (age GT 30 OR age LT 20)"
	tokens, err := lexer.NewLexer(source).Tokenize()
	if err != nil {
		panic(err)
	}

	for _, t := range tokens {
		println(t.PrintToken())
	}
}
