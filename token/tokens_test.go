package token

import (
	"testing"
)

func TestWhitespaceString(t *testing.T) {
	s := Whitespace(2).String()
	if s != "  " {
		t.Error("Unexpected String()", s)
	}
}

func TestWhitespaceTokenString(t *testing.T) {
	s := Whitespace(2).TokenString()
	if s != "token.Whitespace(2)" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestTabString(t *testing.T) {
	s := Tab(2).String()
	if s != "\t\t" {
		t.Error("Unexpected String()", s)
	}
}

func TestTabTokenString(t *testing.T) {
	s := Tab(2).TokenString()
	if s != "token.Tab(2)" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestCommentString(t *testing.T) {
	s := Comment("foo").String()
	if s != "#foo" {
		t.Error("Unexpected String()", s)
	}
}

func TestCommentTokenString(t *testing.T) {
	s := Comment("foo").TokenString()
	if s != "token.Comment(foo)" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestNewLineString(t *testing.T) {
	s := NewLine{}.String()
	if s != "\n" {
		t.Error("Unexpected String()", s)
	}
}

func TestNewLineTokenString(t *testing.T) {
	s := NewLine{}.TokenString()
	if s != "token.NewLine()" {
		t.Error("Unexpected TokenString()", s)
	}
}
