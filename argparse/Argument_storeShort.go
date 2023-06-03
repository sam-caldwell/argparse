package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/valid"
	"strings"
)

// storeShort - return boolean result
func (arg *Argument) storeShort(argument *string) (err error) {
	if err = valid.IsShortArg(argument); err != nil {
		return err
	}

	if *argument == argHelpShort {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg._Short = strings.ToLower(
		strings.TrimSpace(
			strings.TrimLeft(
				strings.TrimSpace(*argument),
				hyphen)))
	return err
}
