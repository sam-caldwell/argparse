package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
)

// Descriptor - An Descriptor Descriptor
type Descriptor struct {
	_Name     string         // positional arg and final name
	_Short    string         // short argument
	_Long     string         // long argument
	_Type     types.ArgTypes // argument type
	_Required bool           // indicates the argument must provide a value
	_Default  interface{}    // provides a default value when not present
	_Help     string         // provides help text
}
