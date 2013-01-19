package model

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
)

// an unexpected token
type UnexpectedToken struct {
	Token Token
}

func (u UnexpectedToken) ItemString() string {
	return fmt.Sprintf("%T(%v)", u, u.Token.TokenString())
}

func (u UnexpectedToken) Equal(i Item) bool {
	if n, ok := i.(UnexpectedToken); ok {
		return u.Token == n.Token
	}
	return false
}
