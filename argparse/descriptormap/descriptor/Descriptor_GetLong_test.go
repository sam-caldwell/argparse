package descriptor

import (
	"testing"
)

// TestDescriptor_GetLong - Test GetLong() and assume no sanitization when directly setting param
func TestDescriptor_GetLong(t *testing.T) {
	var descriptor Descriptor

	const (
		actual   = "--long_arg--"
		expected = "--long_arg--"
	)

	descriptor.long = actual
	if result := descriptor.GetLong(); result != expected {
		t.Fatalf("GetLong() mismatch: %s %s", result, expected)
	}
}
