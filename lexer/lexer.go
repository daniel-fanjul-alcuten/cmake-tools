package parser

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
)

type lexer struct {
	r []rune
	i int
}

func (l *lexer) eof() bool {
	return l.i >= len(l.r)
}

func (l *lexer) curr() rune {
	return l.r[l.i]
}

func (l *lexer) currSafe() (rune, bool) {
	if l.eof() {
		return 0, false
	}
	return l.curr(), true
}

func (l *lexer) adv() {
	l.i++
}

func (l *lexer) read() rune {
	r := l.curr()
	l.adv()
	return r
}

func (l *lexer) next() (rune, bool) {
	l.adv()
	if l.eof() {
		return 0, false
	}
	return l.curr(), true
}

func (l *lexer) advIf(r rune) bool {
	if l.curr() == r {
		l.adv()
		return true
	}
	return false
}

func (l *lexer) advIfSafe(r rune) bool {
	if l.eof() {
		return false
	}
	return l.advIf(r)
}

func (l *lexer) count(c uint, r rune) uint {
	for l.advIfSafe(r) {
		c++
	}
	return c
}

func (l *lexer) lex(tokens chan<- Token) {

L:
	for !l.eof() {
		switch r := l.read(); {

		case r == ' ':
			tokens <- Whitespace(l.count(1, ' '))

		case r == '\t':
			tokens <- Tab(l.count(1, '\t'))

		case r == '(':
			tokens <- OpenParenthesis{}

		case r == ')':
			tokens <- CloseParenthesis{}

		case r == '\n':
			tokens <- NewLine{}

		case r == '#':
			t := ""
			s, ok := l.currSafe()
			for ok && s != '\n' {
				t = t + string(s)
				s, ok = l.next()
			}
			tokens <- Comment(t)

		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
			t := string(r)
			s, ok := l.currSafe()
			for ok && (s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z') {
				t = t + string(s)
				s, ok = l.next()
			}
			tokens <- LiteralString(t)

		case r == '\'':
			t := ""
			s, ok := l.currSafe()
			for ok && s != '\'' {
				t = t + string(s)
				s, ok = l.next()
			}
			if ok {
				l.adv()
				tokens <- SingleQuotedString(t)
			} else {
				tokens <- UnexpectedEof{}
				break L
			}

		case r == '"':
			// TODO deduplicate code
			t := ""
			s, ok := l.currSafe()
			for ok && s != '"' {
				t = t + string(s)
				s, ok = l.next()
			}
			if ok {
				l.adv()
				tokens <- DoubleQuotedString(t)
			} else {
				tokens <- UnexpectedEof{}
				break L
			}

		default:
			tokens <- UnexpectedRune(r)
			break L
		}
	}

	close(tokens)
}

func Lex(str string, tokens chan<- Token) {
	l := lexer{[]rune(str), 0}
	l.lex(tokens)
}
