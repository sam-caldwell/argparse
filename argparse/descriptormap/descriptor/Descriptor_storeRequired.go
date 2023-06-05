package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// storeRequired - store the required flag and type-check the default value
func (arg *descriptor.Descriptor) storeRequired(t types.ArgTypes, required bool, argDefault any) (err error) {
	arg.defaultValue = argDefault
	arg.required = required
	arg.typ = t
	return t.Typecheck(argDefault)
}
