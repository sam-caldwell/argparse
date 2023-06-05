package parsedset

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/parsedelement"
	"github.com/sam-caldwell/argparse/v2/types"
)

// Add - Add a new argument to the parsed argument element map
func (arg *ArgumentElementMap) Add(n string, t types.ArgTypes, v any) error {
	if arg.data == nil {
		arg.data = make(map[string]parsedelement.ArgumentElement)
	}
	if value, ok := arg.data[n]; ok {
		if value.GetType() != t {
			return fmt.Errorf(errTypeCannotChange, n)
		}
		if err := value.SetValue(v); err != nil {
			return err
		}
		arg.data[n] = value
	} else {
		// Create a new ArgumentElement
		var element parsedelement.ArgumentElement

		if err := element.Set(t, v); err != nil {
			return err
		}
		// Store the new record in the map
		arg.data[n] = element
	}
	return nil
}
