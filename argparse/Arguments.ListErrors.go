package argparse

import "fmt"

// ListErrors - List any errors and return the string list
func (arg *Arguments) ListErrors() (result []string) {
	if arg.HasErrors() {
		fmt.Println(errParsingCliArgs)
		for i := 0; i < arg.err.Count(); i++ {
			err := arg.err.Peek(i)
			if err != nil {
				result = append(result, err.Error())
			}
		}
	}
	return result
}
