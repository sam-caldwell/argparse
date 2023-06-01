package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"testing"
)

func TestPositionalArguments(t *testing.T) {
	var arg PositionalArguments
	arg.argName = "test"
	arg.argType = types.String
	arg.argDefault = int64(0)
	arg.argHelp = "testHelp"
}
