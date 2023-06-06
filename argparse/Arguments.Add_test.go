package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArguments_Add(t *testing.T) {
	var arg Arguments

	t.Log("Counting baseline 0 records")
	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}

	p := 0
	for i := 0; i < 5; i++ {
		t.Logf("adding %d", i)
		for n, required := range []bool{true, false} {
			name := fmt.Sprintf("name%d.%d", i, n)
			short := fmt.Sprintf("-%d", p)
			long := fmt.Sprintf("--arg%d%d", i, n)
			typ := types.Boolean
			value := required
			help := fmt.Sprintf("help string %d%d", i, n)
			arg.Add(name, short, long, typ, required, value, help)
			p++
		}
	}
	t.Log("test records added")

	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}
	t.Log("count check passed")

	if count := arg.ErrorCount(); count != 0 {
		t.Log("error count has errors")
		for i, e := range arg.err.List() {
			t.Log(i, e)
		}
		t.Fatalf("Expected no errors.  Got: %d", count)
	}
	t.Log("we have no errors")

	if arg.descriptors.Get("name0.0").GetShort() != "-0" {
		t.Fatalf("Expected name0.0 to have -0 short arg")
	}

	if arg.descriptors.Get("name0.0").GetLong() != "--arg00" {
		t.Fatalf("Expected name0.0 to have -0 short arg")
	}

	if arg.descriptors.Get("name0.0").GetHelp() != "help string 00" {
		t.Fatalf("Expected name0.0 to have -0 short arg")
	}
}
