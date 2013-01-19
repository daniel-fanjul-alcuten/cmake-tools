package rule

import (
	. "github.com/daniel-fanjul-alcuten/cmake-tools/lexer"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/model"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/parser"
	. "github.com/daniel-fanjul-alcuten/cmake-tools/token"
	"testing"
)

var message = "there are tabs"

func TestNoTabsRuleCheck(t *testing.T) {
	assertChecks(t, "#foo")
}

func TestNoTabsRuleCheckWithSpaces(t *testing.T) {
	assertChecks(t, " #foo")
}

func TestNoTabsRuleCheckWithTabs(t *testing.T) {
	assertChecks(t, "\t#foo", message)
}

func TestNoTabsRuleCheckWithSpacesAndTabs(t *testing.T) {
	assertChecks(t, " \t#foo", message)
}

func assertChecks(t *testing.T, str string, expected ...string) {

	tokens := make(chan Token, 8)
	go Lex(str, tokens)

	items := make(chan Item, 4)
	go Parse(tokens, items)

	errs := make(chan Error, 4)
	go NoTabsRule{}.Check(items, errs)

	i := 0
	for error := range errs {
		if i < len(expected) {
			if error.Error() != expected[i] {
				t.Error("Different error", error, ",", expected[i], "at", i)
			}
		} else {
			t.Error("Unexpected error", error, "at", i)
		}
		i++
	}
	for i < len(expected) {
		t.Error("Missing error", expected[i], "at", i)
		i++
	}
}

func TestNoTabsRuleFormat(t *testing.T) {
	assertFormats(t, "#foo", "#foo")
}

func TestNoTabsRuleFormatWithSpaces(t *testing.T) {
	assertFormats(t, " #foo", " #foo")
}

func TestNoTabsRuleFormatWithTabs(t *testing.T) {
	assertFormats(t, "\t#foo", "  #foo")
}

func TestNoTabsRuleFormatWithSpacesAndTabs(t *testing.T) {
	assertFormats(t, " \t#foo", "   #foo")
}

func assertFormats(t *testing.T, str string, expected string) {

	tokens := make(chan Token, 8)
	go Lex(str, tokens)

	items := make(chan Item, 4)
	go Parse(tokens, items)

	formats := make(chan Item, 4)
	go NoTabsRule{}.Format(items, formats)

	i := 0
	for item := range formats {
		if i < 1 {
			if item.String() != expected {
				t.Error("Different Item", item.String(), ",", expected, "at", i)
			}
		} else {
			t.Error("Unexpected Item", item.String(), "at", i)
		}
		i++
	}
	if i < 1 {
		t.Error("Missing Item", expected, "at", i)
		i++
	}
}
