package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"testing"
)

func TestSearchStringList(t *testing.T) {
	expected := []string{
		"a", "b", "c", "d", "e",
	}
	testFunc := func(value string, expectFind bool) {
		arg := types.ArgString(value)
		result, found := searchStringList(&arg, &expected)

		if expectFind {
			if found {
				if result == value {
					t.Logf("OK  : result (%s) expected (%v) found (%v)", result, expectFind, found)
					return
				}
				t.Fatalf("FAIL: result (%s) expected (%v) found (%v)", result, expectFind, found)
			}
		}
		if found {
			t.Fatalf("FAIL: result (%s) found but not expected", result)
		} else {
			t.Logf("OK  : result (%s) not found (not expected to be found)", result)
		}

	}

	testFunc("a", true)
	testFunc("b", true)
	testFunc("c", true)
	testFunc("d", true)
	testFunc("e", true)
	testFunc("f", false)
	testFunc("g", false)

}
