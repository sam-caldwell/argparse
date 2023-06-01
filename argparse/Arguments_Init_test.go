package argparse

import (
	"os"
	"path"
	"testing"
)

func TestArguments_Init(t *testing.T) {
	var args Arguments
	var programName string = path.Base(os.Args[0])
	args.Init("TestDescription")
	if args.description != "TestDescription" {
		t.Fatal("Expected description mismatch")
	}
	if args.programName == os.Args[0] {
		t.Fatal("Program name is expected to have been popped")
	}
	if args.programName != programName {
		t.Fatal("Program name mismatch")
	}
	if args.optional == nil {
		t.Fatal("args.optional is nil (error)")
	}
	if args.value == nil {
		t.Fatal("args.value is nil (error)")
	}
}
