package main

import (
	"fmt"
	"monolang/tokenizer"
)

func main() {
	program := "print 52 print 68 invalid value print 20 print 23.1"
	tokens := tokenizer.Tokenize(program)

	fmt.Println(program)
	fmt.Println(tokens)
}
