package argument

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/types"
)

// storeRequired - store the required flag and type-check the default value
func (arg *Descriptor) storeRequired(t types.ArgTypes, required bool, argDefault any) (err error) {
	arg.defaultValue = argDefault
	if required && (argDefault == nil) {
		return fmt.Errorf(errMissingDefault, t.String())
	}
	arg.required = required
	return t.Typecheck(argDefault)
}
