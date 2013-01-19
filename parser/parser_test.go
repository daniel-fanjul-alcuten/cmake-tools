package parser

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/lexer"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"testing"
)

func TestParseEmptyNil(t *testing.T) {
	assertItems(t, "")
}

func TestParseEmptyWhitespaces(t *testing.T) {
	assertItems(t, "  ", Empty{Whitabs: []Token{Whitespace(2)}})
}

func TestParseEmptyTabs(t *testing.T) {
	assertItems(t, "\t\t", Empty{Whitabs: []Token{Tab(2)}})
}

func TestParseEmptyComment(t *testing.T) {
	comment := Comment("foo")
	assertItems(t, "#foo", Empty{Comment: &comment})
}

func TestParseEmptyNewLine(t *testing.T) {
	assertItems(t, "\n", Empty{Newline: &NewLine{}})
}

func TestParseEmptyFull(t *testing.T) {
	comment := Comment("foo")
	assertItems(t, " \t #foo\n", Empty{
		Whitabs: []Token{Whitespace(1), Tab(1), Whitespace(1)},
		Comment: &comment,
		Newline: &NewLine{}})
}

func assertItems(t *testing.T, str string, expected ...Item) {

	tokens := make(chan Token, 8)
	go Lex(str, tokens)

	items := make(chan Item, 4)
	go Parse(tokens, items)

	i := 0
	for item := range items {
		if i < len(expected) {
			if !item.Equal(expected[i]) {
				t.Error("Different item", item.ItemString(), ",", expected[i].ItemString(), "at", i)
			}
		} else {
			t.Error("Unexpected item", item.ItemString(), "at", i)
		}
		i++
	}
	for i < len(expected) {
		t.Error("Missing item", expected[i].ItemString(), "at", i)
		i++
	}
}
