package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/types"
)

// Get - Fetch a named argument and return value or error
func (args *Arguments) Get(name string) (value interface{}, err error) {
	this, ok := args.optional[types.ArgString(name)]
	if !ok {
		err = fmt.Errorf(errMissingArgument, name)
	}
	value = args.value[types.ArgString(name)]
	if value == nil {
		if this.argRequired && (this.argDefault == nil) {
			return nil, fmt.Errorf(errMissingRequiredArg, name)
		}
		value = this.argDefault
	}
	switch this.argType {
	case types.Boolean:
		fallthrough
	case types.Flag:
		value = value.(bool)
	case types.Float:
		value = value.(float64)
	case types.Integer:
		value = value.(int64)
	case types.String:
		value = value.(string)
	default:
		value = nil
		err = fmt.Errorf(errInternalError)
	}
	return value, err
}
