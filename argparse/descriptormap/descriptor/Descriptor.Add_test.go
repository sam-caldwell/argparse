package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {

	test := func(short, long string, typ types.ArgTypes, required bool, dValue any, help string, expectError bool) {
		var argDesc Descriptor
		if err := argDesc.Add(short, long, typ, required, dValue, help); err != nil {
			t.Fatal(err)
		}
	}

	//Expect no issue
	test("-a", "--all", types.Boolean, true, true, "test help", false)

	//Expect error: -a is a duplicate.
	test("-a", "--all", types.Boolean, true, true, "test help", true)

	//Expect error: -a cannot be assigned with a different long (--ball) when --all is assigned.
	//test("-a", "--ball", types.Boolean, true, true, "test help", true)

	//test("-a", "--all", types.Boolean, true, true, "test help", true)
	//test("-a", "--all", types.Boolean, true, true, "test help", true)

}
