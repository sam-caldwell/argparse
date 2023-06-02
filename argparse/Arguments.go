package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// Arguments - Argument Parser state
type Arguments struct {
	debug       bool
	programName string
	descriptor  []string
	positional  []PositionalArguments
	optional    map[types.ArgString]OptionalParameters
	value       map[types.ArgString]interface{}
}
