package descriptormap

import "fmt"

// ListPositionalArgs - return a formatted string list of positional arguments (name, type, help string)
func (m *Map) ListPositionalArgs(format string) (result []any) {

	for name, argument := range m.data {
		if (argument.GetShort() == "") || (argument.GetLong() == "") {
			typ := argument.GetType()
			result = append(result, fmt.Sprintf(format, name, typ.String(), argument.GetHelp()))
		}
	}
	return result
}
