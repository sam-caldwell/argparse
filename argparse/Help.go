package argparse

import (
	"fmt"
	"strings"
)

// Help - Print Argument usage (help text)
func (args *Arguments) Help() (text string) {
	text = fmt.Sprintf("\n%s\n", args.programName)
	if strings.TrimSpace(args.description) != "" {
		text += "\n" + args.description + "\n"
	}
	if strings.TrimSpace(args.copyright) != "" {
		text += "\n" + args.copyright + "\n"
	}
	for argName, argument := range args.optional {
		text += fmt.Sprintf("\t%s [%s] - %s\n",
			argName,
			argument.argType.String(),
			argument.argHelp)
	}
	return text
}
