package valid

import (
	"regexp"
)

// IsShortArg - return true if *argument is a valid short argument (-{char})
func IsShortArg(argument *string) bool {
	return regexp.MustCompile(validShortArgRegex).MatchString(*argument)
}
