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
		if e, ok := item.(Empty); ok {
			if e.Whitabs != nil {
				for _, whitab := range e.Whitabs {
					_, tab := whitab.(Tab)
					if tab {
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
	// TODO format
}
