package parsedset

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsedelement"
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestArgumentElementMap_Lookup(t *testing.T) {
	var set ArgumentElementMap
	set.data = make(map[string]parsedelement.ArgumentElement)

	test := func(n string, typ types.ArgTypes, value any, expectRecord bool) {
		var element parsedelement.ArgumentElement
		if err := element.Set(typ, value); err != nil {
			t.Fatal(err)
		}
		set.data[n] = element
		searchResult := set.Lookup(&n)

		if searchResult == nil {
			t.Fatalf("expected record but encountered none (%s)", n)
		} else {
			return
		}
	}
	test("firstArg", types.Boolean, true, true)
	test("firstArg", types.Boolean, false, true)
	test("firstArg", types.Integer, 0, true)
}
