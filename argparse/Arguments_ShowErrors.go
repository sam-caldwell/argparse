package argparse

import "fmt"

// ShowErrors - print the list of errors to stdout
func (arg *Arguments) ShowErrors() *Arguments {
	if arg.HasErrors() {
		fmt.Println("Error parsing command-line arguments")
		for _, err := range arg.err {
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}
		}
	} else {
		fmt.Println("No errors in command-line arguments.")
	}
	return arg
}
