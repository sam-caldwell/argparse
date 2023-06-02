package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// AddArgument - Add an argument configuration.
func (args *Arguments) AddArgument(name string, argType types.ArgTypes, required bool,
	argDefault interface{}, helpStr string) (err error) {
	argName := types.ArgString(strings.TrimSpace(strings.ToLower(name)))
	if args.optional == nil {
		return fmt.Errorf(errNotInitialized)
	}
	args.Debug("Arguments object is initialized")
	if (argName == shortHelp) || (argName == longHelp) {
		return fmt.Errorf(errReservedArg, shortHelp, longHelp)
	}
	args.Debug("Arg is not short or long ", argName)
	if !argType.Valid() {
		return fmt.Errorf(errInvalidArgType, argType.String())
	}
	args.Debug("Arg Type is valid ", argName)
	if isValidArgument(&argName) {
		args.optional[types.ArgString(name)] = OptionalParameters{
			argType:     argType,
			argDefault:  argDefault,
			argHelp:     types.ArgHelp(helpStr),
			argRequired: required,
		}
		args.Debug("Argument is a valid optional argument")
	} else {
		if isPositional(&name) {
			args.positional = append(args.positional, PositionalArguments{
				argName:    argName,
				argType:    argType,
				argDefault: argDefault,
			})
			args.Debug("Argument is a valid positional argument")
		} else {
			return fmt.Errorf(errInvalidArgument, name)
		}
	}
	args.Debug("AddArgument() is done.  Error:", err)
	return err
}
