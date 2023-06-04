package argument

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/valid"
	"strings"
)

// storeLong - store the positional (named) argument
func (arg *Descriptor) storeName(argument *string) (err error) {
	if err = valid.IsNameArg(argument); err != nil {
		return err
	}

	if *argument == argHelpName {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg._Name = strings.ToLower(
		strings.TrimSpace(
			strings.TrimLeft(
				strings.TrimSpace(*argument),
				hyphen)))

	return err
}
