package parsedelement

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestArgumentElement_SetType(t *testing.T) {
	var input ArgumentElement
	expectPass := func(typ types.ArgTypes) {
		if err := input.SetType(typ); err != nil {
			t.Fatal(err)
		}
		if o := input.GetType(); typ != o {
			t.Fatalf("Expected Type not set: %s", o.String())
		}
	}
	expectFail := func(typ types.ArgTypes) {
		if err := input.SetType(typ); err == nil {
			t.Fatal(err)
		}
	}

	expectPass(types.Boolean)
	expectPass(types.Flag)
	expectPass(types.Float)
	expectPass(types.Integer)
	expectPass(types.String)
	expectFail(types.ArgTypes(99))
}
