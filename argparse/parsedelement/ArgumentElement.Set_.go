package parsedelement

import "github.com/sam-caldwell/argparse/v2/types"

// Set  - Set the given types.ArgTypes and value for an ArgumentElement
func (arg *ArgumentElement) Set(t types.ArgTypes, v any) (err error) {
	if err := t.Typecheck(v); err != nil {
		return err
	}
	arg.typ = t
	arg.value = v
	return nil
}
