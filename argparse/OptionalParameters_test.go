package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"testing"
)

func TestOptionalParameters(t *testing.T) {
	var arg OptionalParameters
	arg.argType = types.String
	arg.argDefault = int64(0)
	arg.argHelp = "testHelp"
}
