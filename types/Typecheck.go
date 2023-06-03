package types

import "fmt"

// Typecheck -type-check the default value
func (arg *ArgTypes) Typecheck(argDefault any) (err error) {
	var ok bool
	if argDefault == nil {
		err = nil
	} else {
		switch argDefault {
		case Boolean:
			if _, ok = argDefault.(bool); ok {
				return fmt.Errorf(eMsgTypeCheckBoolean)
			}
		case Flag:
			if _, ok = argDefault.(bool); ok {
				return fmt.Errorf(eMsgTypeCheckFlag)
			}
		case Float:
			if _, ok = argDefault.(float64); ok {
				return fmt.Errorf(eMsgTypeCheckFloat)
			}
		case Integer:
			if _, ok = argDefault.(int64); ok {
				return fmt.Errorf(eMsgTypeCheckInteger)
			}
		case String:
			if _, ok = argDefault.(string); ok {
				return fmt.Errorf(eMsgTypeCheckString)
			}
		default:
			return fmt.Errorf(eMsgTypeCheckUnknown)
		}
	}
	return err
}
