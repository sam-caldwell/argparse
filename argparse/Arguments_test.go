package argparse

import (
	"testing"
)

func TestArguments_Struct(t *testing.T) {
	var arg Arguments
	if arg.descriptors != nil {
		t.Fail()
	}
	if arg.programName != "" {
		t.Fail()
	}
	if arg.preambleText != nil {
		t.Fail()
	}
	if arg.postscriptText != nil {
		t.Fail()
	}
	if arg.err.Count() != 0 {
		t.Fail()
	}
	if arg.results.Count() != 0 {
		t.Fail()
	}

}
