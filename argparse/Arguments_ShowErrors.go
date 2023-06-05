package argparse

import "fmt"

// ShowErrors - print the list of errors to stdout
func (arg *Arguments) ShowErrors() *Arguments {
	if arg.HasErrors() {
		fmt.Println(errParsingCliArgs)
		for i := 0; i < arg.err.Count(); i++ {
			err := arg.err.Peek(i)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	}
	return arg
}
