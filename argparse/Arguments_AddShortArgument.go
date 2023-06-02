package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
	"strings"
)

// AddShortArgument - add a short argument (-<char>)
func (args *Arguments) AddShortArgument(
	shortArg string,
	name string,
	argType types.ArgTypes,
	required bool,
	argDefault interface{},
	helpStr string) (err error) {

	argName := strings.TrimSpace(strings.TrimSpace(name))

	if !isShortArgument(&shortArg) {
		return fmt.Errorf("expected short argument for %s", argName)
	}

	return args.addOptionalArgument(
		&argName,
		argType,
		required,
		&argDefault,
		&helpStr)
}
