package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/valid"
	"strings"
)

// storeHelpString - sanitize help string.  Store if valid.  return error state.
func (arg *Argument) storeHelpString(help *string) (err error) {
	if err = valid.IsValidHelpString(help); err != nil {
		return err
	}
	arg._Help = strings.TrimSpace(*help)
	return nil
}
