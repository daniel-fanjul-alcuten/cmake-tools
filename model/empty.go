package model

import (
	"fmt"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
)

// something that is not a Command
type Empty struct {
	// Whitespace's and Tab's
	Whitabs []Token
	Comment *Comment
	Newline *NewLine
}

func (e Empty) ItemString() string {
	s := ""
	if e.Whitabs != nil {
		for _, t := range e.Whitabs {
			if len(s) > 0 {
				s += ", "
			}
			s += t.TokenString()
		}
	}
	if e.Comment != nil {
		if len(s) > 0 {
			s += ", "
		}
		s += e.Comment.TokenString()
	}
	if e.Newline != nil {
		if len(s) > 0 {
			s += ", "
		}
		s += e.Newline.TokenString()
	}
	return fmt.Sprintf("%T(%v)", e, s)
}

func (e Empty) Equal(i Item) bool {
	if m, ok := i.(Empty); ok {
		if e.Whitabs == nil {
			if m.Whitabs != nil {
				return false
			}
		} else if m.Whitabs == nil {
			return false
		} else if len(e.Whitabs) != len(m.Whitabs) {
			return false
		} else {
			for i, token := range e.Whitabs {
				if token != m.Whitabs[i] {
					return false
				}
			}
		}
		if e.Comment == nil {
			if m.Comment != nil {
				return false
			}
		} else if m.Comment == nil {
			return false
		} else if *e.Comment != *m.Comment {
			return false
		}
		if e.Newline == nil {
			if m.Newline != nil {
				return false
			}
		} else if m.Newline == nil {
			return false
		} else if *e.Newline != *m.Newline {
			return false
		}
		return true
	}
	return false
}
