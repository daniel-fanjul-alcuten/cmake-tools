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
	"os"
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

			first_rule_items := make(chan Item, 1024)
			last_rule_items := first_rule_items
			for _, rule := range rules {
				next_rule_items := make(chan Item, 1024)
				go rule.Format(last_rule_items, next_rule_items)
				last_rule_items = next_rule_items
			}

			write := true

			go func() {
				str := ""
				for item := range last_rule_items {
					str += item.String()
				}
				if write {
					if err := ioutil.WriteFile(filename+".tmp", []byte(str), 0666); err != nil {
						log.Println(err)
					} else {
						if err := os.Rename(filename+".tmp", filename); err != nil {
							log.Println(err)
						}
					}
				}
				wait <- true
			}()

			for item := range items {
				if ut, ok := item.(UnexpectedToken); ok {
					log.Printf("%v: unexpected token: %v", filename, ut.Token.TokenString())
					write = false
				}
				if write {
					first_rule_items <- item
				}
			}
			close(first_rule_items)
		}()
		nwaits += 1
	}

	for nwaits > 0 {
		<-wait
		nwaits--
	}
}
