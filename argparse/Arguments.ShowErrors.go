package argparse

import "fmt"

// ShowErrors - print the list of errors to stdout
func (arg *Arguments) ShowErrors() *Arguments {
	if arg.err.Count() > 0 {
		fmt.Println(errParsingCliArgs)

		for _, line := range arg.Errors() {
			fmt.Printf("%s\n", line)
		}
	}
	return arg
}
