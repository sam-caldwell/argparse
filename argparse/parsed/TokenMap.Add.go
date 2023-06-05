package parsed

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed/token"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Add - Add a new token to the TokenMap
func (arg *TokenMap) Add(name *string, typ types.ArgTypes, value *any) error {
	if arg.data == nil {
		arg.data = make(map[string]token.Token)
	}
	if thisArg, ok := arg.data[*name]; ok {
		if thisArg.GetType() != typ {
			return fmt.Errorf(errTypeCannotChange, *name)
		}
		if err := thisArg.SetValue(*value); err != nil {
			return err
		}
		arg.data[*name] = thisArg
	} else {
		// Create a new Token
		var element token.Token

		if err := element.Set(typ, *value); err != nil {
			return err
		}
		// Store the new record in the map
		arg.data[*name] = element
	}
	return nil
}
