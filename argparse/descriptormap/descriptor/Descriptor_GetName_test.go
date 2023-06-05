package descriptor

import (
	"testing"
)

func TestDescriptor_GetName(t *testing.T) {
	var descriptor Descriptor
	const original = "This_Is_A_Test"
	descriptor.name = original
	if result := descriptor.GetName(); result != "This_Is_A_Test" {
		t.Fatalf("actual/expected mismatch %s %s", original, result)
	}
}
