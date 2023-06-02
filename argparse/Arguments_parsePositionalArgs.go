package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
)

// parsePositionalArgs - parse the positional arguments
func (args *Arguments) parsePositionalArgs() (err error) {

	args.Debugf("os.Arg, positional args count: %d < %d", len(os.Args), len(args.positional))

	// We expect cli args to be at least the number of positional args.
	if len(os.Args) < len(args.positional) {
		return fmt.Errorf(errNotEnoughArgs)
	}

	args.Debug("positional arg count ok")

	// Loop through the positional arguments and each value in os.Args
	// is the value we need.  We then pop this value off the os.Args
	// list and what will remain when we are done are optional only.
	for pos := 0; pos < len(args.positional); pos++ {
		args.Debugf("pos: %d", pos)
		argName := args.positional[pos].argName
		args.Debugf("argName: %s", argName)
		value, err := pop(&os.Args)
		if err != nil {
			args.Debugf("pop() returned err: %s", err)
			return err
		}
		switch args.positional[pos].argType {
		case types.Boolean:
			args.value[argName] = value
		case types.Flag:
			args.value[argName] = value
		case types.Float:
			args.value[argName] = value
		case types.Integer:
			args.value[argName] = value
		case types.String:
			args.value[argName] = value
		default:
			args.Debugf("argName (%s) type (%s) mismatch", argName, args.positional[pos].argType.String())
			return fmt.Errorf(errTypeMismatch,
				args.positional[pos].argType.String(),
				argName)
		}
	}
	defer args.Debugf("parsePositionalArgs() return.  err:%v", err)
	return err
}
