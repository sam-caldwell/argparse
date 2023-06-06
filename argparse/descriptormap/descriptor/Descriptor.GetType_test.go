package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestDescriptor_GetType(t *testing.T) {
	var arg Descriptor
	arg.typ = types.String

	if typ := arg.GetType(); typ != "String" {
		t.Fatalf("Descriptor.GetType() failed.  Got: %s", typ)
	}
}
