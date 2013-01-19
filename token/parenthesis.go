package token

import (
	"fmt"
)

// an open parenthesis
type OpenParenthesis struct{}

func (t OpenParenthesis) String() string {
	return "("
}

func (t OpenParenthesis) TokenString() string {
	return fmt.Sprintf("%T()", t)
}

// close parenthesis
type CloseParenthesis struct{}

func (t CloseParenthesis) String() string {
	return ")"
}

func (t CloseParenthesis) TokenString() string {
	return fmt.Sprintf("%T()", t)
}
