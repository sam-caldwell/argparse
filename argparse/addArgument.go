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
	if !isValidArgument(&argName) {
		return fmt.Errorf(errInvalidArgument, name)
	}
	if (argName == shortHelp) || (argName == longHelp) {
		return fmt.Errorf(errReservedArg, shortHelp, longHelp)
	}
	if args.optional == nil {
		return fmt.Errorf(errNotInitialized)
	}
	if !argType.Valid() {
		return fmt.Errorf(errInvalidArgType, argType.String())
	}
	if isPositional(&name) {
		args.positional = append(args.positional, PositionalArguments{
			argName:    argName,
			argType:    argType,
			argDefault: argDefault,
		})
	} else {
		args.optional[types.ArgString(name)] = OptionalParameters{
			argType:     argType,
			argDefault:  argDefault,
			argHelp:     types.ArgHelp(helpStr),
			argRequired: required,
		}
	}
	return err
}
