package descriptor

import "testing"

func TestDescriptorMap(t *testing.T) {
	var d DescriptorMap
	if d != nil {
		t.Fatal("initial state should be nil")
	}
	d = make(DescriptorMap)
	d["test1"] = Descriptor{}
	if len(d) != 1 {
		t.Fatal("expected size 1(a)")
	}
	d["test1"] = Descriptor{}
	if len(d) != 1 {
		t.Fatal("expected size 1(b)")
	}
	d["test2"] = Descriptor{}
	if len(d) != 2 {
		t.Fatal("expected size 2")
	}
}
