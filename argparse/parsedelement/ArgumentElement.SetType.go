package parsedelement

import "github.com/sam-caldwell/argparse/v2/types"

// SetType  - Set the given types.ArgTypes value for an ArgumentElement
func (arg *ArgumentElement) SetType(t types.ArgTypes) (err error) {
	if err := t.Valid(); err != nil {
		return err
	}
	arg.typ = t
	return nil
}
