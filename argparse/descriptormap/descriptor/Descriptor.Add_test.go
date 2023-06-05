package descriptor

import (
	"errors"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {

	test := func(short, long string, typ types.ArgTypes, required bool, dValue any, help string, expectError error) {
		var argDesc Descriptor
		if err := argDesc.Add(short, long, typ, required, dValue, help); err != nil {
			if expectError != nil {
				if err != expectError {
					t.Fatal(err)
				}
			} else {
				t.Fatal(err)
			}
		}
	}

	//Expect no issue
	test("-a", "--all", types.Boolean, true, true, "test help", errors.New(errArgExists))

	//Expect error: -a is a duplicate.
	test("-a", "--all", types.Boolean, true, true, "test help", errors.New(errArgExists))

	//Expect error: -a cannot be assigned with a different long (--ball) when --all is assigned.
	//test("-a", "--ball", types.Boolean, true, true, "test help", true)

	//test("-a", "--all", types.Boolean, true, true, "test help", true)
	//test("-a", "--all", types.Boolean, true, true, "test help", true)

}
