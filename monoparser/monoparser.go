package monoparser

import (
	"monolang/tokenizer"
)

type NodeType string

const (
	PrintStatement NodeType = "printStatement"
)

type ExpressionType string

const (
	NumberLiteral ExpressionType = "numberLiteral"
)

type Expression struct {
	Type  ExpressionType
	Value string
}

type Node struct {
	Type       NodeType
	Expression Expression
}

type AST []Node

func Parse(tokens []tokenizer.Token) AST {
	i := 0
	nodes := make([]Node, 0)

	for i < len(tokens) {
		e, c := parseStatement(i, tokens[i], tokens)
		if c {
			nodes = append(nodes, e)
			i += 2
		}
	}

	return nodes
}

func parseExpression(t tokenizer.Token) (e Expression) {
	switch t.Type {
	case tokenizer.Number:
		e.Type = NumberLiteral
		e.Value = t.Value
	}

	return e
}

func parseStatement(i int, t tokenizer.Token, tokens []tokenizer.Token) (n Node, consumed bool) {
	if t.Type == tokenizer.Keyword {
		switch t.Value {
		case "print":
			consumed = true
			n.Type = PrintStatement
			n.Expression = parseExpression(tokens[i+1])
		}
	}

	return n, consumed
}
