package parser

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"testing"
)

func TestEmpty(t *testing.T) {
	assertTokens(t, "")
}

func TestWhitespaces(t *testing.T) {
	assertTokens(t, "  ", Whitespace(2))
}

func TestTabs(t *testing.T) {
	assertTokens(t, "\t\t", Tab(2))
}

func TestOpenParenthesis(t *testing.T) {
	assertTokens(t, "(", OpenParenthesis{})
}

func TestCloseParenthesis(t *testing.T) {
	assertTokens(t, ")", CloseParenthesis{})
}

func TestNewLine(t *testing.T) {
	assertTokens(t, "\n\n", NewLine{}, NewLine{})
}

func TestComments(t *testing.T) {
	assertTokens(t, "#foo", Comment("foo"))
}

func TestCommentsWithNewLine(t *testing.T) {
	assertTokens(t, "#foo\n", Comment("foo"), NewLine{})
}

func TestLiteralString(t *testing.T) {
	assertTokens(t, "foo", LiteralString("foo"))
}

func TestLiteralStringWithSpaces(t *testing.T) {
	assertTokens(t, "foo ", LiteralString("foo"), Whitespace(1))
}

func TestLiteralStringWithTab(t *testing.T) {
	assertTokens(t, "foo\t", LiteralString("foo"), Tab(1))
}

func TestLiteralStringWithNewLine(t *testing.T) {
	assertTokens(t, "foo\n", LiteralString("foo"), NewLine{})
}

func TestLiteralStringWithComment(t *testing.T) {
	assertTokens(t, "foo#bar", LiteralString("foo"), Comment("bar"))
}

func TestSingleQuotedString(t *testing.T) {
	assertTokens(t, "'foo'", SingleQuotedString("foo"))
}

func TestSingleQuotedStringUnfinishedEmpty(t *testing.T) {
	assertTokens(t, "'", UnexpectedEof{})
}

func TestSingleQuotedStringUnfinishedNotEmpty(t *testing.T) {
	assertTokens(t, "'foo", UnexpectedEof{})
}

func TestSingleQuotedStringWithSpaces(t *testing.T) {
	assertTokens(t, "'foo' ", SingleQuotedString("foo"), Whitespace(1))
}

func TestSingleQuotedStringWithTab(t *testing.T) {
	assertTokens(t, "'foo'\t", SingleQuotedString("foo"), Tab(1))
}

func TestSingleQuotedStringWithNewLine(t *testing.T) {
	assertTokens(t, "'foo'\n", SingleQuotedString("foo"), NewLine{})
}

func TestSingleQuotedStringWithComment(t *testing.T) {
	assertTokens(t, "'foo'#bar", SingleQuotedString("foo"), Comment("bar"))
}

func TestDoubleQuotedString(t *testing.T) {
	assertTokens(t, "\"foo\"", DoubleQuotedString("foo"))
}

func TestDoubleQuotedStringUnfinishedEmpty(t *testing.T) {
	assertTokens(t, "\"", UnexpectedEof{})
}

func TestDoubleQuotedStringUnfinishedNotEmpty(t *testing.T) {
	assertTokens(t, "\"foo", UnexpectedEof{})
}

func TestDoubleQuotedStringWithSpaces(t *testing.T) {
	assertTokens(t, "\"foo\" ", DoubleQuotedString("foo"), Whitespace(1))
}

func TestDoubleQuotedStringWithTab(t *testing.T) {
	assertTokens(t, "\"foo\"\t", DoubleQuotedString("foo"), Tab(1))
}

func TestDoubleQuotedStringWithNewLine(t *testing.T) {
	assertTokens(t, "\"foo\"\n", DoubleQuotedString("foo"), NewLine{})
}

func TestDoubleQuotedStringWithComment(t *testing.T) {
	assertTokens(t, "\"foo\"#bar", DoubleQuotedString("foo"), Comment("bar"))
}

func assertTokens(t *testing.T, str string, expected ...Token) {

	tokens := make(chan Token, 8)
	go Lex(str, tokens)

	i := 0
	for token := range tokens {
		if i < len(expected) {
			if token != expected[i] {
				t.Error("Different token", token.TokenString(), ",", expected[i].TokenString(), "at", i)
			}
		} else {
			t.Error("Unexpected token", token.TokenString(), "at", i)
		}
		i++
	}
	for i < len(expected) {
		t.Error("Missing token", expected[i].TokenString(), "at", i)
		i++
	}
}
