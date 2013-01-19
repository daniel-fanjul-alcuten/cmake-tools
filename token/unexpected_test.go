package token

import (
	"testing"
)

func TestUnexpectedRuneString(t *testing.T) {
	s := UnexpectedRune('f').String()
	if s != "token.UnexpectedRune('f')" {
		t.Error("Unexpected String()", s)
	}
}

func TestUnexpectedRuneTokenString(t *testing.T) {
	s := UnexpectedRune('f').TokenString()
	if s != "token.UnexpectedRune('f')" {
		t.Error("Unexpected TokenString()", s)
	}
}

func TestUnexpectedEofString(t *testing.T) {
	s := UnexpectedEof{}.String()
	if s != "token.UnexpectedEof()" {
		t.Error("Unexpected String()", s)
	}
}

func TestUnexpectedEofTokenString(t *testing.T) {
	s := UnexpectedEof{}.TokenString()
	if s != "token.UnexpectedEof()" {
		t.Error("Unexpected TokenString()", s)
	}
}
