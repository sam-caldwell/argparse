package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
)

// Descriptor - An Descriptor Descriptor
type Descriptor struct {
	_Short    string
	_Long     string
	_Name     string
	_Type     types.ArgTypes
	_Required bool
	_Default  interface{}
	_Help     string
}
