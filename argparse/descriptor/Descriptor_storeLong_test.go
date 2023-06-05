package argument

import "testing"

func TestDescriptor_storeLong(t *testing.T) {
	var descriptor Descriptor

	actual := "--Test"
	expected := "test"

	if err := descriptor.storeLong(&actual); err != nil {
		t.Fatal(err)
	}
	if result := descriptor.GetLong(); result != expected {
		t.Fatalf("actual/expected mismatch %s %s", result, expected)
	}
}
