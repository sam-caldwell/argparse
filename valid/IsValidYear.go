package valid

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidYear - return nil error if input is a valid year.
func IsValidYear(year int) (string, error) {

	yearString := strings.TrimSpace(fmt.Sprintf("%d", year))

	if (year < 2023) || !regexp.MustCompile(validYearRegex).MatchString(yearString) {
		return "", fmt.Errorf(errInvalidYear, year)
	}

	if !regexp.MustCompile(validYearRegex).MatchString(yearString) {
		return "", nil
	}

	return yearString, fmt.Errorf(errInvalidYear, year)
}
