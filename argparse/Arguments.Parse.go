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

	for position, rawToken := range os.Args[1:] {

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

		//Deal with the positional arguments (token is a value now)
		if optionalArgEncountered {
			//If we have already encountered a short -a or long --all optional argument, bail.
			_ = arg.err.Push(fmt.Errorf(errPositionalArgumentFollowingOptional))
			return arg
		}

		//Otherwise look up our expected argument by position.
		thisName, thisDescriptor = arg.descriptors.GetByPosition(position)

		//But if the result (thisName) is nil, we weren't expecting an arg at this position.  bail.
		if thisName == nil {
			// if position not found in our descriptor map then we have a syntax error.
			_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
			return arg
		}

		//However, if all this works out, token must be a positional value.  Type-check it. Bail on fail.
		typ := thisDescriptor.GetType()
		if err := typ.Typecheck(token); err != nil {
			_ = arg.err.Push(err)
			return arg
		}

		// Store this value
		var value any = token
		if err := arg.results.Add(thisName, typ, &value); err != nil {
			panic(err)
		}

		// Set our state to expect a new token
		valueExpected = false

		continue
	} /*end of for*/
	_ = arg.err.Push(fmt.Errorf(errParsingCliArgs, *thisName))
	return arg
}
