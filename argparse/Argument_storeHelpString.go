package argparse

import (
	"fmt"
	"strings"
)

// storeHelpString - store a given help string
func (arg *Argument) storeHelpString(help *string) (err error) {
	if help == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	arg._Help = strings.TrimSpace(*help)
	return nil
}
