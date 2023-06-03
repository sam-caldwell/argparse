package valid

import (
	"regexp"
)

func IsShortArg(argument *string) bool {
	return regexp.MustCompile(validShortArgRegex).MatchString(*argument)
}
