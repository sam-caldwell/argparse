package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"testing"
)

func TestArguments(t *testing.T) {
	var arg Arguments
	arg.programName = ""
	arg.copyright = ""
	arg.description = ""
	arg.positional = append(arg.positional, PositionalArguments{
		argName:    "test",
		argType:    types.String,
		argDefault: int64(0),
		argHelp:    "testHelp",
	})
	arg.optional = make(map[types.ArgString]OptionalParameters, 1)
	arg.optional["test"] = OptionalParameters{
		argType:     types.Integer,
		argDefault:  float32(0.0),
		argHelp:     "testHelp",
		argRequired: true,
	}
	arg.value = make(map[types.ArgString]interface{})
	arg.value["testValue"] = "myString"
}
