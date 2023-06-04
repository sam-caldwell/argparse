package argument

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/valid"
	"strings"
)

// storeShort - store a short arg -{char}
func (arg *Descriptor) storeShort(argument *string) (err error) {
	if err = valid.IsShortArg(argument); err != nil {
		return err
	}

	if *argument == argHelpShort {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg.short = strings.ToLower(
		strings.TrimSpace(
			strings.TrimLeft(
				strings.TrimSpace(*argument),
				hyphen)))
	return err
}