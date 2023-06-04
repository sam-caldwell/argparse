package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

// TestDescriptor_Struct - Test initial state of the Descriptor struct
func TestDescriptor_Struct(t *testing.T) {
	var d Descriptor
	if d.name != "" {
		t.Fail()
	}
	if d.short != "" {
		t.Fail()
	}
	if d.long != "" {
		t.Fail()
	}
	if d.typ != types.String {
		t.Fail()
	}
	if d.required {
		t.Fail()
	}
	if d.defaultValue != nil {
		t.Fail()
	}
	if d.help != "" {
		t.Fail()
	}
}
