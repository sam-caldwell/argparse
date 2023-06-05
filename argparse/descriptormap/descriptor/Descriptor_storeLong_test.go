package descriptor

import (
	"strings"
	"testing"
)

func TestDescriptor_storeLong(t *testing.T) {
	test := func(arg string, expectedError bool) {
		var descriptor Descriptor
		if err := descriptor.storeLong(&arg); err != nil {
			if !expectedError {
				t.Fatal(err)
			}
		}
		expected := strings.ToLower(
			strings.TrimSpace(
				strings.TrimLeft(
					strings.TrimSpace(arg),
					hyphen)))

		if strings.TrimLeft(arg, hyphen) != expected {
			t.Fatalf("long mismatch (%s)(%s)", arg, expected)
		}

	}
	test("-a", true)
	test("", true)
	test("--arg", false)
	test("invalid", true)

}
