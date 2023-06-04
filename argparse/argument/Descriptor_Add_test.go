package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {
	var descriptor Descriptor

	err := descriptor.Add("argName", "-a", "--arg", types.Integer,
		true, 1337, "help string")

	if err != nil {
		t.Fatal(err)
	}

	if descriptor.name != "argName" {
		t.Fail()
	}

	if descriptor.short != "-a" {
		t.Fail()
	}

	if descriptor.long != "--arg" {
		t.Fail()
	}

	if descriptor.typ != types.Integer {
		t.Fail()
	}

	if !descriptor.required {
		t.Fail()
	}

	if descriptor.defaultValue != 1337 {
		t.Fail()
	}

	if descriptor.help != "help string" {
		t.Fail()
	}
}
