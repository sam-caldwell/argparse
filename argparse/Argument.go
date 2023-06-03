package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// Argument - An Argument Descriptor
type Argument struct {
	_Short    string
	_Long     string
	_Name     string
	_Type     types.ArgTypes
	_Required bool
	_Default  interface{}
	_Help     string
}
