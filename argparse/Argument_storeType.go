package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// storeType - return boolean result
func (arg *Argument) storeType(argType types.ArgTypes) (err error) {
	if argType.Valid() {
		return nil
	}
	return fmt.Errorf(errInvalidArgType, argType.String())
}
