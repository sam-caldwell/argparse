package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"testing"
)

func TestDescriptorMap(t *testing.T) {
	var d Map
	if d != nil {
		t.Fatal("initial state should be nil")
	}
	d = make(Map)
	d["test1"] = descriptor.Descriptor{}
	if len(d) != 1 {
		t.Fatal("expected size 1(a)")
	}
	d["test1"] = descriptor.Descriptor{}
	if len(d) != 1 {
		t.Fatal("expected size 1(b)")
	}
	d["test2"] = descriptor.Descriptor{}
	if len(d) != 2 {
		t.Fatal("expected size 2")
	}
}
