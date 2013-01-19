package rule

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
)

type Error struct {
	// TODO add line numbers
	error
}

type Rule interface {
	Check(items <-chan Item, errs chan<- Error)
	Format(input <-chan Item, output chan<- Item)
}
