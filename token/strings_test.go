package token

import (
	"testing"
)

func TestLiteralStringString(t *testing.T) {
	s := LiteralString("foo").String()
	if s != "foo" {
		t.Error("Unexpected String()", s)
	}
}

func TestLiteralStringTokenString(t *testing.T) {
	s := LiteralString("foo").TokenString()
	if s != "token.LiteralString(foo)" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestSingleQuotedStringString(t *testing.T) {
	s := SingleQuotedString("foo").String()
	if s != "'foo'" {
		t.Error("Unexpected String()", s)
	}
}

func TestSingleQuotedStringTokenString(t *testing.T) {
	s := SingleQuotedString("foo").TokenString()
	if s != "token.SingleQuotedString(foo)" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestDoubleQuotedStringString(t *testing.T) {
	s := DoubleQuotedString("foo").String()
	if s != "\"foo\"" {
		t.Error("Unexpected String()", s)
	}
}

func TestDoubleQuotedStringTokenString(t *testing.T) {
	s := DoubleQuotedString("foo").TokenString()
	if s != "token.DoubleQuotedString(foo)" {
		t.Error("Unexpected TokenString()", s)
	}
}
