package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Add - append the given Argument object to the list of Arguments.
func (arg *Arguments) Add(
	name string,
	short string,
	long string,
	typ types.ArgTypes,
	required bool,
	argDefault any,
	help string) *Arguments {

	/*
	 * +-----------+----------+------------------+-------------------------------------------- - - - -
	 * | name       | string   | required         | must be unique and non-empty string.
	 * | short      | string   | "" | -{char}     | if blank, short arg will not be evaluated
	 * | long       | string   | "" | --{string}  | if blank, long arg will not be evaluated
	 * | typ        | required | String           | argument data type (Boolean, Flag, Float, Integer, String
	 * | required   | bool     | default=false    |
	 * | argDefault | any      | nil | any        | optional (must eval to type of `typ`)
	 * | help       | string   | arbitrary string | help string
	 * +-----------+----------+------------------+-------------------------------------------- - - - -
	 */

	// Create a new descriptor
	var thisDescriptor descriptor.Descriptor

	// Add the state to the descriptor and push any error if not nil
	_ = arg.err.Push(thisDescriptor.Add(short, long, typ, required, argDefault, help))

	return arg
}
