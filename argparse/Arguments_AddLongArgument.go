package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// AddLongArgument - Add a long argument (--string)
func (args *Arguments) AddLongArgument(
	name string,
	argType types.ArgTypes,
	required bool,
	argDefault interface{},
	helpStr string) (err error) {

	argName := strings.TrimSpace(strings.TrimLeft(strings.ToLower(name), "-"))

	if !isLongArgument(&argName) {
		return fmt.Errorf(errExpectedLongArg, argName)
	}

	return args.addOptionalArgument(
		&argName,
		argType,
		required,
		&argDefault,
		&helpStr)
}
