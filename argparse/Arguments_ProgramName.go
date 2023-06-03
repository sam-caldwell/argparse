package argparse

import "os"

// ProgramName - Set the program name from os.Args[0]
func (arg *Arguments) ProgramName() *Arguments {
	arg.programName = os.Args[0]
	return arg
}
