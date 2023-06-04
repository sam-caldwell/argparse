package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
	"strings"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {
	var descriptor Descriptor

	const (
		name       = "argName"
		short      = "-a"
		long       = "--arg"
		typ        = types.Integer
		required   = true
		defaultArg = int64(1337)
		help       = "help string"
	)

	err := descriptor.Add(name, short, long, typ, required, defaultArg, help)
	if err != nil {
		t.Fatal(err)
	}

	if descriptor.name != strings.ToLower(name) {
		t.Fatal(descriptor.name)
	}

	if descriptor.short != strings.ToLower(strings.TrimLeft(short, "-")) {
		t.Fatal(descriptor.short)
	}

	if descriptor.long != strings.ToLower(strings.TrimLeft(long, "-")) {
		t.Fatal(descriptor.long)
	}

	if descriptor.typ != types.Integer {
		t.Fatal(descriptor.typ.String())
	}

	if !descriptor.required {
		t.Fatal(descriptor.required)
	}

	if descriptor.defaultValue != defaultArg {
		t.Fatal(descriptor.defaultValue)
	}

	if descriptor.help != help {
		t.Fatal(descriptor.help)
	}
}
