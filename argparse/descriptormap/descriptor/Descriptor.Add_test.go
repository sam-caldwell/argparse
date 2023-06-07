package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/counters/v2"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {

	test := func(pos *counters.ConditionalCounter, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		var argDesc Descriptor

		if err := argDesc.Add(pos, short, long, typ, required, dValue, help); err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	//Expect no issue
	var pos counters.ConditionalCounter
	test(&pos, "-a", "--all", types.Boolean, true, true, "test help")

	//Expect error: -a is a duplicate.
	test(&pos, "-a", "--all", types.Boolean, true, true, "test help")

	test(&pos, "", "", types.Boolean, true, true, "test help")

}
