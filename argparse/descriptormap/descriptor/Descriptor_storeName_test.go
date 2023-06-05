package descriptor

import (
	"testing"
)

func TestDescriptor_storeName(t *testing.T) {
	var descriptor Descriptor

	actual := "This_Is_A_Test"
	expected := "this_is_a_test"

	if err := descriptor.storeName(&actual); err != nil {
		t.Fatal(err)
	}
	if result := descriptor.GetName(); result != expected {
		t.Fatalf("actual/expected mismatch %s %s", result, expected)
	}
}
