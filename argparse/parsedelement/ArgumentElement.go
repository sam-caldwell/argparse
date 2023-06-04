package parsedelement

import "github.com/sam-caldwell/argparse/v2/types"

// ArgumentElement - this is a parsed argument from command-line
type ArgumentElement struct {
	typ   types.ArgTypes
	value any
}
