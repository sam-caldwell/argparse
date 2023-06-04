package types

import (
	"fmt"
)

// Typecheck -type-check the default value
func (arg *ArgTypes) Typecheck(argDefault any) (err error) {
	var ok bool
	if argDefault != nil {
		switch *arg {
		case Boolean:
			_, ok = argDefault.(bool)
		case Flag:
			_, ok = argDefault.(bool)
		case Float:
			_, ok = argDefault.(float64)
		case Integer:
			_, ok = argDefault.(int)
		case String:
			_, ok = argDefault.(string)
		default:
			return fmt.Errorf(eMsgTypeCheckUnknown)
		}
		if !ok {
			return fmt.Errorf(eMsgTypeCheck, (*arg).String())
		}
	}
	return err
}
