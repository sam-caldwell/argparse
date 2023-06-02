package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// addOptionalArgument - Add an optional argument object (short or long) after sanitization to the Arguments object.
func (args *Arguments) addOptionalArgument(
	argName *string,
	argType types.ArgTypes,
	required bool,
	argDefault *interface{},
	helpStr *string) (err error) {

	// check for an uninitialized Arguments object
	if args.optional == nil {
		return fmt.Errorf(errNotInitialized)
	}

	if !argType.Valid() {
		return fmt.Errorf(errInvalidArgType, argType.String())
	}

	if reservedShortArgument(argName) || reservedLongArgument(argName) {
		return fmt.Errorf(errReservedArg, argHelpShort, argHelpLong)
	}

	args.optional[types.ArgString(*argName)] = OptionalParameters{
		argType:     argType,
		argDefault:  *argDefault,
		argHelp:     types.ArgHelp(strings.TrimSpace(*helpStr)),
		argRequired: required,
	}

	return nil
}
