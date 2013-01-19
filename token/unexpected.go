package token

import (
	"fmt"
)

// unexpected token
type UnexpectedRune rune

func (t UnexpectedRune) String() string {
	return t.TokenString()
}

func (t UnexpectedRune) TokenString() string {
	return fmt.Sprintf("%T(%q)", t, uint(t))
}

// unexpected eof
type UnexpectedEof struct{}

func (t UnexpectedEof) String() string {
	return t.TokenString()
}

func (t UnexpectedEof) TokenString() string {
	return fmt.Sprintf("%T()", t)
}
