package token

import (
	"fmt"
	"strings"
)

// lexical token
type Token interface {
	String() string
	TokenString() string
}

// number of consecutive whitespaces
type Whitespace uint

func (t Whitespace) String() string {
	return strings.Repeat(" ", int(t))
}

func (t Whitespace) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, uint(t))
}

// number of consecutive tabs
type Tab uint

func (t Tab) String() string {
	return strings.Repeat("\t", int(t))
}

func (t Tab) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, uint(t))
}

// a comment without the #
type Comment string

func (t Comment) String() string {
	return "#" + string(t)
}

func (t Comment) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, string(t))
}

// a new line
type NewLine struct{}

func (t NewLine) String() string {
	return "\n"
}

func (t NewLine) TokenString() string {
	return fmt.Sprintf("%T()", t)
}
