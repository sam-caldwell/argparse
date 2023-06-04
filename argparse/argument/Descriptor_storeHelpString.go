package argument

import (
	"github.com/sam-caldwell/argparse/v2/valid"
	"strings"
)

// storeHelpString - sanitize help string.  Store if valid.  return error state.
func (arg *Descriptor) storeHelpString(help *string) (err error) {
	if err = valid.IsValidHelpString(help); err != nil {
		return err
	}
	arg._Help = strings.TrimSpace(*help)
	return nil
}
