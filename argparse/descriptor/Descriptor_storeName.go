package argument

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/valid"
	"strings"
)

// storeName - store the positional (named) argument
func (arg *Descriptor) storeName(argument *string) (err error) {

	if err = valid.IsNameArg(argument); err != nil {
		return err
	}

	if *argument == argHelpName {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg.name = strings.ToLower(strings.TrimSpace(*argument))

	return err
}
