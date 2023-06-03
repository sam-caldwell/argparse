package valid

import (
	"fmt"
	"regexp"
)

// IsNameArg - return nil error if *argument is a valid positional argument name
func IsNameArg(argument *string) error {
	if !regexp.MustCompile(validArgRegex).MatchString(*argument) {
		return nil
	}
	return fmt.Errorf(errExpectedNameArg, *argument)
}
