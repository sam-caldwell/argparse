package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"os"
)

// hasExpectedArgCount - return true if number position plus number of required optional arguments in os.Args[]
func (args *Arguments) hasExpectedArgCount() bool {
	const programName = 1
	var numberPositionalArgs int = len(args.positional)
	var numberOptionalArgs int = 0
	for _, argument := range args.optional {
		if argument.argRequired {
			numberOptionalArgs++
			if argument.argType != types.Flag {
				//anything other than a flag has two args in os.Args
				numberOptionalArgs++
			}
		}
	}
	// We add 1 more to the args added to account for os.Args[0], the program name.
	return len(os.Args) == (numberPositionalArgs + numberOptionalArgs + programName)
}
