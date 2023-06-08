package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"strings"
	"testing"
)

const expected = `
argparse.test
  This is a test program
  It does things...really cool things.

Usage:
  argparse.test [positional args] [optional args]

  Positional Arguments
    noshort [String] - test no
    pos [Integer] - positional

 Optional Arguments
    -h           --help                          - show this help message
    -t [Boolean] --test [Boolean] [default:true] - Test 1                
    -a [Boolean] --all  [Boolean] [default:true] - Test all              
    -n           --now                           - test now              
                 --nos  [String]  [default:ns]   - test no               

 This program has a postscript
 The postscript comes after usage.
 (c) 2023 Sam Caldwell <mail@samcaldwell.net>


`

func TestArguments_Help(t *testing.T) {
	var arg Arguments
	arg.ProgramName().
		Preamble("This is a test program").
		Preamble("It does things...really cool things.").
		Postscript("This program has a postscript").
		Postscript("The postscript comes after usage.").
		Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net").
		Add("pos", "", "", types.Integer, true, 1, "positional").
		Add("test", "-t", "--test", types.Boolean, true, true, "Test 1").
		Add("all", "-a", "--all", types.Boolean, true, true, "Test all").
		Add("now", "-n", "--now", types.String, true, "notnow", "test now").
		Add("now", "-n", "--now", types.Flag, true, false, "test now").
		Add("noshort", "", "--nos", types.String, true, "ns", "test no")

	helpText := arg.Help()

	if arg.HasErrors() {
		t.Log("Errors:")
		for _, e := range arg.Errors() {
			t.Log(e)
		}
		t.Log("----")
		t.Fatal("Failing due to errors")
	}

	actualText := strings.TrimSpace(helpText)
	expectedText := strings.TrimSpace(expected)

	fmt.Println("---")
	fmt.Println(actualText)
	fmt.Println("---")

	if actualText != expectedText {
		actualLines := strings.Split(actualText, "\n")
		expectedLines := strings.Split(expectedText, "\n")

		if len(actualLines) != len(expectedLines) {
			t.Fatalf("count mismatch (actual: %d, expected: %d", len(actualLines), len(expectedLines))
		}

		for line, _ := range expectedLines {
			matches := actualLines[line] == expectedLines[line]
			if !matches {
				szActual := len(actualLines[line])
				szExpected := len(expectedLines[line])

				if szActual != szExpected {
					t.Logf("  actual:%2d :'%s'", line, actualLines[line])
					t.Logf("expected:%2d :'%s'", line, expectedLines[line])
					t.Fatalf("%d: actual (%d) not expected (%d)", line, szActual, szExpected)
				}
				t.Logf("  actual:%2d :'%s'", line, actualLines[line])
				t.Logf("expected:%2d :'%s'", line, expectedLines[line])
			}
		}
		t.Fatal("help text does not match expected text")
	}
}
