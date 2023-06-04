package parsed

import "github.com/sam-caldwell/argparse/v2/types"

// ArgumentElement - this is a parsed argument from command-line
type ArgumentElement struct {
	_name  string
	_type  types.ArgTypes
	_value any
}
