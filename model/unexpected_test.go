package model

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"testing"
)

func TestUnexpectedTokenString(t *testing.T) {
	s := UnexpectedToken{Whitespace(1)}.String()
	if s != "model.UnexpectedToken(token.Whitespace(1))" {
		t.Error("Unexpected String()", s)
	}
}

func TestUnexpectedTokenItemString(t *testing.T) {
	s := UnexpectedToken{Whitespace(1)}.ItemString()
	if s != "model.UnexpectedToken(token.Whitespace(1))" {
		t.Error("Unexpected ItemString()", s)
	}
}
