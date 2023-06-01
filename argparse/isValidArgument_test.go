package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"testing"
)

func TestIsValidArgument(t *testing.T) {
	testFunc := func(value types.ArgString, expected bool) {
		result := isValidArgument(&value)
		t.Logf("result: %v", result)
		if result != expected {
			t.Fatalf("expected '%s' argument to return %v. expected %v", value, result, expected)
		}
	}
	testFunc("-i", true)
	testFunc("--valid", true)
	testFunc("--valid_arg", true)
	testFunc("--valid-arg", true)
	testFunc("invalid", false)
	testFunc("1nvalid", false)
	testFunc("invalid_", false)
	testFunc("invalid@", false)
}
