package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType string

const (
	Number     TokenType = "number"
	Keyword    TokenType = "keyword"
	Whitespace TokenType = "whitespace"
)

type Token struct {
	Index int
	Type  TokenType
	Value string
	Valid bool
}

type MatcherFunc func(input string, index int) Token

var Keywords [1]string = [1]string{"print"}
var KeywordRegEx string = fmt.Sprintf("^(%s)", strings.Join(Keywords[:], "|"))

var Matchers [3]MatcherFunc = [3]MatcherFunc{
	createRegexMatcher("^[.0-9]+", Number),
	createRegexMatcher(KeywordRegEx, Keyword),
	createRegexMatcher("^\\s+", Whitespace),
}

func Tokenize(input string) []Token {
	tokens := make([]Token, 0)
	i := 0
	for i < len(input) {
		matches := make([]Token, 0)
		for _, matcher := range Matchers {
			match := matcher(input, i)
			if match.Valid {
				match.Index = i
				matches = append(matches, match)
			}
		}
		if len(matches) > 0 {
			match := matches[0]
			tokens = append(tokens, match)
			i = i + len(match.Value)
		}

		i++
	}

	return tokens
}

func createRegexMatcher(regex string, t TokenType) MatcherFunc {
	return func(input string, index int) Token {
		tok := Token{
			Type:  t,
			Valid: true,
		}

		substring := input[index:]
		re := regexp.MustCompile(regex)
		match := re.FindString(string(substring))

		if len(match) > 0 && t != Whitespace {
			tok.Value = string(match)
		} else {
			tok.Valid = false
		}

		return tok
	}
}
