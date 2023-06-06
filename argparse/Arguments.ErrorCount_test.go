package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArguments_ErrorCount(t *testing.T) {
	var arg Arguments

	t.Log("Counting baseline 0 records")
	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}

	t.Log("Adding Invalid record")
	arg.Add("", "", "", types.ArgTypes(99), true, nil, "BadRecord")

	t.Log("Checking error count 0 (should be 1)")
	if count := arg.ErrorCount(); count != 0 {
		t.Fatal("Expected error did not happen")
	}

	t.Log("Checking error count 1 (should be 1)")
	if count := arg.ErrorCount(); count != 1 {
		t.Fatal("Expected error count should be 1")
	}
}
