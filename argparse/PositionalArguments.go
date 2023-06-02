package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// PositionalArguments - This struct represents a positional argument struct
type PositionalArguments struct {
	argName    types.ArgString
	argType    types.ArgTypes
	argDefault interface{}
	argHelp    types.ArgHelp
}
