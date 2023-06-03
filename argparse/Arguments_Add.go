package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// Add - append the given Argument object to the list of Arguments.
func (arg *Arguments) Add(name string, short string, long string,
	argType types.ArgTypes, required bool, argDefault any,
	help string) *Arguments {

	var newArgument Argument
	err := newArgument.Add(name, short, long, argType, required, argDefault, help)
	if err != nil {
		arg.err = err
	}
	(*arg).set = append((*arg).set, newArgument)
	return arg
}
