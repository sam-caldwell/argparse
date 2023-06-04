package argument

import "testing"

// TestStoreShort - test Descriptor.storeShort()
func TestStoreShort(t *testing.T) {
	var descriptor Descriptor
	expected := "-a"
	if err := descriptor.storeShort(&expected); err != nil {
		t.Fatal(err)
	}
	if descriptor.short != expected {
		t.Fatal("short descriptor mismatch")
	}
}
