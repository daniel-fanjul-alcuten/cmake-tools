package parser

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
)

func Parse(tokens <-chan Token, items chan<- Item) {

	token, ok := <-tokens
	for ok {
		whitabs := getWhitabs(tokens, &token, &ok)
		comment := getComment(tokens, &token, &ok)
		newline := getNewLine(tokens, &token, &ok)
		if whitabs != nil || comment != nil || newline != nil {
			items <- Empty{whitabs, comment, newline}
		} else {
			items <- UnexpectedToken{token}
			break
		}
	}
	for ok {
		_, ok = <-tokens
	}
	close(items)
}

func getWhitabs(tokens <-chan Token, token *Token, ok *bool) []Token {
	whitabs := []Token(nil)
	for *ok {
		switch (*token).(type) {
		case Whitespace, Tab:
			if whitabs == nil {
				whitabs = make([]Token, 0, 2)
			}
			whitabs = append(whitabs, *token)
			*token, *ok = <-tokens
		default:
			return whitabs
		}
	}
	return whitabs
}

func getComment(tokens <-chan Token, token *Token, ok *bool) *Comment {
	if *ok {
		if comment, okc := (*token).(Comment); okc {
			*token, *ok = <-tokens
			return &comment
		}
	}
	return nil
}

func getNewLine(tokens <-chan Token, token *Token, ok *bool) *NewLine {
	if *ok {
		if newline, oknl := (*token).(NewLine); oknl {
			*token, *ok = <-tokens
			return &newline
		}
	}
	return nil
}
