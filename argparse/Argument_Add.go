package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// Add - Sanitize an argument object's values
func (arg *Argument) Add(name string, short string, long string,
	argType types.ArgTypes, required bool, argDefault any,
	help string) (err error) {

	if err = arg.storeName(&name); err != nil {
		return err
	}

	if err = arg.storeShort(&short); err != nil {
		return err
	}

	if err = arg.storeLong(&long); err != nil {
		return err
	}

	if err = arg.storeType(argType); err != nil {
		return err
	}

	if err = arg.storeRequired(argType, required, argDefault); err != nil {
		return err
	}

	if err = arg.storeHelpString(&help); err != nil {
		return err
	}

	return nil

}
