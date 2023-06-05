package argparse

import (
	"fmt"
	"os"
)

func (arg *Arguments) ExitOnError() *Arguments {
	if arg.HasErrors() {
		fmt.Println("Error parsing command-line arguments")
		arg.ShowErrors()
		fmt.Println("")
		fmt.Println(arg.Help())
		os.Exit(exitArgParseError)
	}
	return arg
}
