package argument

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/valid"
	"strings"
)

// storeLong - store a long argument (--string)
func (arg *Descriptor) storeLong(argument *string) (err error) {
	if err = valid.IsLongArg(argument); err != nil {
		return err
	}

	if *argument == argHelpLong {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg.long = strings.ToLower(
		strings.TrimSpace(
			strings.TrimLeft(
				strings.TrimSpace(*argument),
				hyphen)))
	return err
}
