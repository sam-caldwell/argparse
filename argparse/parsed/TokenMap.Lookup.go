package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsedelement"
)

// Lookup - return object if found or nil of not found
func (arg *Argument) Lookup(n *string) *parsedelement.ArgumentElement {
	if arg.data != nil {
		if thisArg, found := (*arg).data[*n]; found {
			return &thisArg
		}
	}
	return nil
}
