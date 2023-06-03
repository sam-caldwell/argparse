package valid

import "regexp"

func IsLongArg(argument *string) bool {
	return regexp.MustCompile(validLongArgRegex).MatchString(*argument)
}
