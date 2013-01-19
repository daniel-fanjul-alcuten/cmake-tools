package main

import (
	"flag"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/lexer"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/parser"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/rule"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	rules := []Rule{NoTabsRule{}}

	nwaits, wait := 0, make(chan bool, 128)
	for _, arg := range args {
		filename := arg

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println(err)
			continue
		}

		tokens := make(chan Token, 10*1024)
		go Lex(string(data), tokens)

		items := make(chan Item, 1024)
		go Parse(tokens, items)

		go func() {

			rule_items := make([]chan Item, 0, len(rules))
			for _, rule := range rules {
				r_items := make(chan Item, 1024)
				r_errors := make(chan Error, 256)
				go rule.Check(r_items, r_errors)
				go func() {
					for error := range r_errors {
						log.Printf("%v: %v", filename, error)
					}
					wait <- true
				}()
				rule_items = append(rule_items, r_items)
			}

			for item := range items {
				if ut, ok := item.(UnexpectedToken); ok {
					log.Printf("%v: unexpected token: %v", filename, ut.Token.TokenString())
				}
				for i := range rules {
					rule_items[i] <- item
				}
			}

			for i := range rules {
				close(rule_items[i])
			}
		}()
		nwaits += len(rules)
	}

	for nwaits > 0 {
		<-wait
		nwaits--
	}
}
