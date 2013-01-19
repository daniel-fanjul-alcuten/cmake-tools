package model

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"testing"
)

func TestEmptyItemString(t *testing.T) {
	comment := Comment("foo")
	s := Empty{
		Whitabs: []Token{Whitespace(1), Tab(1), Whitespace(1)},
		Comment: &comment,
		Newline: &NewLine{}}.ItemString()
	if s != "model.Empty(token.Whitespace(1), token.Tab(1), token.Whitespace(1), token.Comment(foo), token.NewLine())" {
		t.Error("Unexpected ItemString()", s)
	}
}
