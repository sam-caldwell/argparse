package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// Arguments - Argument Parser state
type Arguments struct {
	programName string
	copyright   string
	description string
	positional  []PositionalArguments
	optional    map[types.ArgString]OptionalParameters
	value       map[types.ArgString]interface{}
}
