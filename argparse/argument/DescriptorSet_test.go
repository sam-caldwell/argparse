package argument

import "testing"

func TestDescriptorMap(t *testing.T) {
	var d DescriptorMap
	if d != nil {
		t.Fatal("initial state should be nil")
	}
	d = make(DescriptorMap)
	d["test"] = Descriptor{}
	if len(d) != 1 {
		t.Fatal("expected size 1")
	}
}
