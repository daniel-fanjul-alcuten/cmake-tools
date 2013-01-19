package token

import (
	"testing"
)

func TestOpenParenthesisString(t *testing.T) {
	s := OpenParenthesis{}.String()
	if s != "(" {
		t.Error("Unexpected String()", s)
	}
}

func TestOpenParenthesisTokenString(t *testing.T) {
	s := OpenParenthesis{}.TokenString()
	if s != "token.OpenParenthesis()" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestCloseParenthesisString(t *testing.T) {
	s := CloseParenthesis{}.String()
	if s != ")" {
		t.Error("Unexpected String()", s)
	}
}

func TestCloseParenthesisTokenString(t *testing.T) {
	s := CloseParenthesis{}.TokenString()
	if s != "token.CloseParenthesis()" {
		t.Error("Unexpected TokenString()", s)
	}
}
