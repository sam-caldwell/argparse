package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestDescriptor_storeRequired_int64(t *testing.T) {
	var descriptor Descriptor

	argType := types.Integer
	argDefault := int64(1000)

	if err := descriptor.storeRequired(types.Integer, true, argDefault); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != argType {
		t.Fatalf("type mismatch %v %v", descriptor.typ, argType)
	}
	if descriptor.required != true {
		t.Fatal("required mismatch")
	}
	if descriptor.defaultValue.(int64) != argDefault {
		t.Fatal("defaultValue mismatch")
	}
}
