package valid

import (
	"fmt"
	"regexp"
)

// IsLongArg - return nil error if *argument is a valid long argument (--{string})
func IsLongArg(argument *string) error {
	if argument == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	if regexp.MustCompile(validLongArgRegex).MatchString(*argument) {
		return nil
	}
	return fmt.Errorf(errExpectedLongArg, *argument)
}
