package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

type PositionalArguments struct {
	argName    types.ArgString
	argType    types.ArgTypes
	argDefault interface{}
	argHelp    types.ArgHelp
}
