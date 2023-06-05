package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
)

// Descriptor - A Descriptor struct
type Descriptor struct {
	name         string         // positional arg and final name
	short        string         // short argument
	long         string         // long argument
	typ          types.ArgTypes // argument type
	required     bool           // indicates the argument must provide a value
	defaultValue interface{}    // provides a default value when not present
	help         string         // provides help text
}
