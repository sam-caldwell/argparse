package argument

import "testing"

// TestDescriptor_GetShort - Test GetShort
func TestDescriptor_GetShort(t *testing.T) {
	var descriptor Descriptor
	descriptor.short = "-o"
	if descriptor.GetShort() != "-o" {
		t.Fail()
	}
}
