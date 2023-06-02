package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// AddPositionalArgument - Add a positional argument following the last positional argument (if any)
func (args *Arguments) AddPositionalArgument(
	name string,
	argType types.ArgTypes,
	argDefault interface{},
	strHelp string) (err error) {

	// check for an uninitialized Arguments object
	if args.optional == nil {
		return fmt.Errorf(errNotInitialized)
	}

	if !argType.Valid() {
		return fmt.Errorf(errInvalidArgType, argType.String())
	}

	argName := types.ArgString(strings.TrimSpace(strings.ToLower(name)))

	if reservedPositionalArgument(&argName) {
		return
	}

	args.positional = append(args.positional, PositionalArguments{
		argName:    argName,
		argType:    argType,
		argDefault: argDefault,
		argHelp:    types.ArgHelp(strHelp),
	})

	return nil
}
