package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// storeRequired - store the required flag and type-check the default value
func (arg *Argument) storeRequired(t types.ArgTypes, required bool, argDefault any) (err error) {
	if required && (argDefault == nil) {
		return fmt.Errorf(errMissingDefault, t.String())
	}
	arg._Required = required
	arg._Default = argDefault
	return t.Typecheck(argDefault)
}