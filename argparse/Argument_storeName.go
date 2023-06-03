package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/valid"
	"strings"
)

// storeLong - return boolean result
func (arg *Argument) storeName(argument *string) (err error) {
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
