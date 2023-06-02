package argparse

import "fmt"

// Dump - Dump the argument names as a list of strings and their values
func (args *Arguments) Dump() (output []string) {
	columnSize := 0
	for k, _ := range args.value {
		sz := len(k)
		if sz > columnSize {
			columnSize = len(k)
		}
	}
	for k, v := range args.value {
		output = append(output, fmt.Sprintf("%15s: %v", k, v))
	}
	return output
}
