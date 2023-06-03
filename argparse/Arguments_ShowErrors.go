package argparse

import "fmt"

func (arg *Arguments) ShowErrors() *Arguments {
	if arg.hasErrors() {
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
