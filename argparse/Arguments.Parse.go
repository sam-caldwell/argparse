package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/argparse/v2/argparse/valid"
	"os"
	"strings"
)

// Parse - Parse os.Args
func (arg *Arguments) Parse() *Arguments {

	var optionalArgEncountered = false
	var valueExpected = false
	var thisName *string
	var thisDescriptor *descriptor.Descriptor = nil

	for _, rawToken := range os.Args {

		// clean the token any intentional whitespace should be wrapped in quotes.
		token := strings.TrimSpace(rawToken)

		// we expect a value and will not consider -a or --all patterns
		if valueExpected {
			var value any
			typ := thisDescriptor.GetType()
			value = token
			if err := arg.results.Add(thisName, typ, &value); err != nil {
				panic(err)
			}
		}

		// Deal with the long arguments...
		if err := valid.IsLongArg(&token); err == nil {
			thisName, thisDescriptor = arg.descriptors.GetByLong(&token)
			if thisName == nil {
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
				return arg //bail!
			}
			//Flags have no values.  We can just store them.
			if typ := thisDescriptor.GetType(); typ == types.Flag {
				var value any = true
				_ = arg.err.Push(arg.results.Add(thisName, types.Flag, &value))
			}
			valueExpected = true //we need the next item to be a value
			optionalArgEncountered = true
			continue //Just go get the next one
		}

		//Deal with the short arguments
		if err := valid.IsShortArg(&token); err == nil {
			thisName, thisDescriptor = arg.descriptors.GetByShort(&token)
			if thisName == nil {
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
				return arg //bail!
			}
			//Flags have no values.  We can just store them.
			if typ := thisDescriptor.GetType(); typ == types.Flag {
				var value any = true
				_ = arg.err.Push(arg.results.Add(thisName, types.Flag, &value))
			}
			valueExpected = true //we need the next item to be a value
			optionalArgEncountered = true
			continue //Just go get the next one
		}

		//Deal with the positional arguments
		if err := valid.IsNameArg(&token); err == nil {
			if optionalArgEncountered {
				_ = arg.err.Push(fmt.Errorf(errPositionalArgumentFollowingOptional))
			}
			thisName, thisDescriptor = arg.descriptors.GetByName(&token)
			if thisName == nil {
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
				return arg
			}
			//Flags have no values.  We can just store them.
			if typ := thisDescriptor.GetType(); typ == types.Flag {
				var value any = true
				_ = arg.err.Push(arg.results.Add(thisName, types.Flag, &value))
			}
			valueExpected = true //we need the next item to be a value
			continue             //Just go get the next one
		}
		_ = arg.err.Push(fmt.Errorf(errParsingCliArgs, token))
		return arg
	}
	return arg
}
