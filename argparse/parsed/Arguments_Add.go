package parsed

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/types"
)

// Add - Add a parsed argument record to the parsed.Arguments object
func (arg *Arguments) Add(n string, t types.ArgTypes, v any) (err error) {
	// if arg.set is not yet initialized, do so.
	if arg.set == nil {
		arg.set = make(ArgumentElementSet)
	}
	// if we previously saved an argument (n), we can only update it if type does not change.
	if thisArgument, found := arg.set[n]; found && (thisArgument._type != t) {
		return fmt.Errorf(errArgumentsCannotChangeTypes, n, t.String(), v)
	}
	//Add the record
	return arg.set.Add(n, t, v)
}
