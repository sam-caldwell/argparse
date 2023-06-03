package valid

import "regexp"

// IsLongArg - return true if *argument is a valid long argument (--{string})
func IsLongArg(argument *string) bool {
	return regexp.MustCompile(validLongArgRegex).MatchString(*argument)
}
