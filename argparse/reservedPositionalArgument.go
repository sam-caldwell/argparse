package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// reservedPositionalArgument - Return true of 'help' is present.
func reservedPositionalArgument(argument *types.ArgString) bool {
	return *argument == argHelp
}
