package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsed/token"
)

// Lookup - return object if found or nil of not found
func (arg *TokenMap) Lookup(n *string) *token.Token {
	if arg.data != nil {
		if thisArg, found := (*arg).data[*n]; found {
			return &thisArg
		}
	}
	return nil
}
