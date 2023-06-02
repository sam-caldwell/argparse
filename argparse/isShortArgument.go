package argparse

import (
	"regexp"
)

// isShortArgument - Return boolean if short argument
func isShortArgument(argument *string) bool {
	const validArgRegex = `^-[a-z]$`
	return regexp.MustCompile(validArgRegex).MatchString(*argument)
}
