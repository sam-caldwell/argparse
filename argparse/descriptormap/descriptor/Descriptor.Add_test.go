package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {

	test := func(pos int, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		var argDesc Descriptor

		if err := argDesc.Add(pos, short, long, typ, required, dValue, help); err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	//Expect no issue
	test(-1, "-a", "--all", types.Boolean, true, true, "test help")

	//Expect error: -a is a duplicate.
	test(-1, "-a", "--all", types.Boolean, true, true, "test help")

	test(0, "", "", types.Boolean, true, true, "test help")

}
