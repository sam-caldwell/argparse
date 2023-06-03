package argparse

import (
	"fmt"
	"regexp"
	"strings"
)

// Copyright - Append the Copyright string to the preamble.
func (arg *Arguments) Copyright(year int, author string, email string) *Arguments {
	const validYearRegex = `^[0-9]{4}$`

	yearString := strings.TrimSpace(fmt.Sprintf("%d", year))
	if (year < 2023) || !regexp.MustCompile(validYearRegex).MatchString(yearString) {
		arg.err = fmt.Errorf("invalid year")
	}

	author = strings.TrimSpace(author)

	email = strings.TrimSpace(email)

	return arg.Preamble(
		fmt.Sprintf("(c) %s %s <%s>",
			yearString,
			author,
			email))
}
