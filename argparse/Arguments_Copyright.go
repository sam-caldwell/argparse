package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-argparse/v2/valid"
	"strings"
)

// Copyright - Append the Copyright string to the preamble.
func (arg *Arguments) Copyright(year int, author string, email string) *Arguments {

	var err error

	yearString, err := valid.IsValidYear(year)
	if err != nil {
		arg.err.Push(err)
	}

	author = strings.TrimSpace(author)

	email = strings.TrimSpace(email)

	return arg.Preamble(
		fmt.Sprintf("(c) %s %s <%s>",
			yearString,
			author,
			email))
}
