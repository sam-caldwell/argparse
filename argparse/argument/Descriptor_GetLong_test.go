package argument

import "testing"

// TestDescriptor_GetLong - Test GetLong()
func TestDescriptor_GetLong(t *testing.T) {
	var descriptor Descriptor
	descriptor.long = "--long_arg"
	if descriptor.GetLong() != "--long_arg" {
		t.Fail()
	}
}
