package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// AddLongArgument - Add an argument configuration.
func (args *Arguments) AddLongArgument(
	name string,
	argType types.ArgTypes,
	required bool,
	argDefault interface{},
	helpStr string) (err error) {

	// check for an uninitialized Arguments object
	if args.optional == nil {
		return fmt.Errorf(errNotInitialized)
	}
	argName := types.ArgString(strings.TrimSpace(strings.TrimLeft(strings.ToLower(name), "-")))

	if reservedArgument(&argName) {
		return fmt.Errorf(errReservedArg, shortHelp, longHelp)
	}
	if isLongArgument(&argName)

	if !argType.Valid() {
		return fmt.Errorf(errInvalidArgType, argType.String())
	}
	if isValidArgument(&argName) {
		args.optional[types.ArgString(name)] = OptionalParameters{
			argType:     argType,
			argDefault:  argDefault,
			argHelp:     types.ArgHelp(helpStr),
			argRequired: required,
		}
	} else {
		if isPositional(&name) {
			args.positional = append(args.positional, PositionalArguments{
				argName:    argName,
				argType:    argType,
				argDefault: argDefault,
			})
		} else {
			return fmt.Errorf(errInvalidArgument, name)
		}
	}
	return err
}
