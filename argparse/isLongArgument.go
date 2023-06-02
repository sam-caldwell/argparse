package argparse

import (
	"regexp"
)

// isLongArgument - Return boolean if long argument
func isLongArgument(argument *string) bool {
	const validArgRegex = `^--[a-z][a-z0-9_-]*[a-z0-9]$`
	return regexp.MustCompile(validArgRegex).MatchString(*argument)
}
