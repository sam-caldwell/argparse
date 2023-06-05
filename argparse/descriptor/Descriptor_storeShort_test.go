package argument

import (
	"strings"
	"testing"
)

// TestStoreShort - test Descriptor.storeShort()
func TestStoreShort(t *testing.T) {
	var descriptor Descriptor
	expected := "-a"
	if err := descriptor.storeShort(&expected); err != nil {
		t.Fatal(err)
	}
	if descriptor.short != strings.TrimPrefix(expected, "-") {
		t.Fatalf("short descriptor mismatch (%s|%s)", descriptor.short, expected)
	}
}
