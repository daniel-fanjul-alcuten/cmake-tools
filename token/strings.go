package token

import (
	"fmt"
)

// a string without separators
type LiteralString string

func (t LiteralString) String() string {
	return string(t)
}

func (t LiteralString) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, string(t))
}

// a string inside single quotes
type SingleQuotedString string

func (t SingleQuotedString) String() string {
	return "'" + string(t) + "'"
}

func (t SingleQuotedString) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, string(t))
}

// a string inside double quotes
type DoubleQuotedString string

func (t DoubleQuotedString) String() string {
	return "\"" + string(t) + "\""
}

func (t DoubleQuotedString) TokenString() string {
	return fmt.Sprintf("%T(%v)", t, string(t))
}
