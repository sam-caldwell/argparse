package parsedelement

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestArgumentElement_Struct(t *testing.T) {
	var arg ArgumentElement
	if arg.typ != types.String {
		t.Fail()
	}
	if arg.value != nil {
		t.Fail()
	}
}
