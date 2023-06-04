package argument

import (
	"github.com/sam-caldwell/argparse/v2/types"
)

// storeType - sanitize the given type.  Store it if valid.  return error state
func (arg *Descriptor) storeType(argType types.ArgTypes) (err error) {
	if err = argType.Valid(); err != nil {
		return err
	}
	arg._Type = argType
	return nil
}
