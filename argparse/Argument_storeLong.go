package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/valid"
	"strings"
)

// storeLong - return boolean result
func (arg *Argument) storeLong(argument *string) (err error) {
	if err = valid.IsLongArg(argument); err != nil {
		return err
	}

	if *argument == argHelpShort {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg._Long = strings.ToLower(
		strings.TrimSpace(
			strings.TrimLeft(
				strings.TrimSpace(*argument),
				hyphen)))
	return err
}
