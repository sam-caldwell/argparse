package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Add - append the given Argument object to the list of Arguments.
func (arg *Arguments) Add(name string, short string, long string,
	typ types.ArgTypes, required bool, argDefault any,
	help string) *Arguments {
	/*
	 * name       - required
	 * short      - if blank, will not be configured
	 * long       - if blank, will not be configured
	 * typ        - required (ArgTypes)
	 * required   - bool
	 * argDefault - any
	 * help       - help string
	 */

	// Create a new descriptor
	var thisDescriptor descriptor.Descriptor

	// Add the state to the descriptor
	arg.err.Push(
		thisDescriptor.Add(name, short, long, typ, required, argDefault, help))

	// Append the descriptor to the Arguments
	arg.descriptors[name] = thisDescriptor

	return arg
}
