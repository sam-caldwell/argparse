package argparse

import (
	"fmt"
)

// Help - Print Argument usage (help text)
func (args *Arguments) Help() (text string) {
	text = fmt.Sprintf("\n%s\n\n", args.programName)

	for _, line := range args.descriptor {
		text += line + "\n"
	}
	for argName, argument := range args.optional {
		text += fmt.Sprintf("\t%s [%s] - %s\n",
			argName,
			argument.argType.String(),
			argument.argHelp)
	}
	return text
}
