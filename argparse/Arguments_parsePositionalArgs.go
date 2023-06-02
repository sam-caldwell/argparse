package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
	"strconv"
)

// parsePositionalArgs - parse the positional arguments
func (args *Arguments) parsePositionalArgs() (err error) {

	args.Debugf("os.Arg, positional args count: %d < %d", len(os.Args), len(args.positional))

	// We expect cli args to be at least the number of positional args.
	if !args.hasExpectedArgCount() {
		return fmt.Errorf(errNotEnoughArgs)
	}

	//Assumption: pop() will remove program name (arg 0) when .Init() is called.
	for _, argument := range args.positional {
		argName := argument.argName
		thisValue, err := pop(&os.Args)
		if err != nil {
			return err
		}
		switch argument.argType {
		case types.Boolean:
			if args.value[argName], err = strconv.ParseBool(thisValue); err != nil {
				return err
			}
		case types.Float:
			if args.value[argName], err = strconv.ParseFloat(thisValue, 64); err != nil {
				return err
			}
		case types.Integer:
			if args.value[argName], err = strconv.ParseInt(thisValue, 10, 64); err != nil {
				return err
			}
		case types.String:
			args.value[argName] = thisValue
		default:
			return fmt.Errorf(errTypeMismatch, argument.argType.String(), argName)
		}
	}
	return err
}
