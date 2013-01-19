package rule

import (
	"errors"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
)

// no tabs allowed
type NoTabsRule struct{}

func (NoTabsRule) Check(items <-chan Item, errs chan<- Error) {
	for item := range items {
		if empty, ok := item.(Empty); ok {
			if empty.Whitabs != nil {
				for _, whitab := range empty.Whitabs {
					if _, ok := whitab.(Tab); ok {
						errs <- Error{errors.New("there are tabs")}
						break
					}
				}
			}
		}
	}
	close(errs)
}

func (NoTabsRule) Format(input <-chan Item, output chan<- Item) {
	for item := range input {
		if empty, ok := item.(Empty); ok {
			if empty.Whitabs != nil {
				nw := make([]Token, len(empty.Whitabs))
				for i, whitab := range empty.Whitabs {
					if tab, ok := whitab.(Tab); ok {
						nw[i] = Whitespace(2 * tab)
					} else {
						nw[i] = whitab
					}
				}
				empty.Whitabs = nw
			}
			output <- empty
		} else {
			output <- item
		}
	}
	close(output)
}
