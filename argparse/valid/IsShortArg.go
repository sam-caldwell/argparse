package valid

import (
	"fmt"
	"regexp"
)

// IsShortArg - return nil error if *argument is a valid short argument (-{char})
func IsShortArg(argument *string) error {
	if argument == nil {
		return fmt.Errorf(errEmptyOrNilObject)
	}
	if regexp.MustCompile(validShortArgRegex).MatchString(*argument) {
		return nil
	}
	return fmt.Errorf(errExpectedShortArg, *argument)
}
