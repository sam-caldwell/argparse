package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArguments_Count(t *testing.T) {
	var arg Arguments

	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}

	p := 0
	for i := 0; i < 10; i++ {
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
	if count := arg.descriptors.Count(); count != p {
		t.Fatalf("Expected %d records. Got %d", p, count)
	}
}
