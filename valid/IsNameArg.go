package valid

import (
	"regexp"
)

// IsNameArg - return true if *argument is a valid positional argument name
func IsNameArg(argument *string) bool {
	return regexp.MustCompile(validArgRegex).MatchString(*argument)
}
