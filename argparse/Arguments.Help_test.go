package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"strings"
	"testing"
)

const expected = `argparse.test
  This is a test program
  It does things...really cool things.

Usage:
  argparse.test [positional args] [optional args]

  Positional Arguments
   pos - positional

  Optional Arguments
  -h                 --help                 (help   ) [            ] - show this help message
  -t <Boolean>       --test <Boolean>       (test   ) [default:true] - Test 1
  -a <Boolean>       --all  <Boolean>       (all    ) [default:true] - Test all
  -n                 --now                  (now    ) [            ] - test now
     <String>        --nos  <String>        (noshort) [default:ns  ] - test no

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

	fmt.Println(strings.TrimSpace(helpText))

	if strings.TrimSpace(helpText) != strings.TrimSpace(expected) {
		actualLines := strings.Split(strings.TrimSpace(helpText), "\n")
		expectedLines := strings.Split(strings.TrimSpace(expected), "\n")

		if len(actualLines) != len(expectedLines) {
			t.Logf("count mismatch (actual: %d, expected: %d", len(actualLines), len(expectedLines))
		}

		for lineNo, line := range expectedLines {
			matches := actualLines[lineNo] == expectedLines[lineNo]
			t.Logf("%d (%v) : %s", lineNo, matches, line)
		}
		t.Fatal("help text does not match expected text")
	}
}
