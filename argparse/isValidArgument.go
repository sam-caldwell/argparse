package argparse

import (
	"github.com/sam-caldwell/go-argparse/v2/types"
	"regexp"
)

// isValidArgument - return boolean if the argument is valid
func isValidArgument(argument *types.ArgString) bool {
	if argument == nil {
		return false
	}

	// Define the regular expression pattern for validating the argument
	const validArgumentRegex = `^(-[a-z]|--[a-z]([a-z0-9_-]*[a-z0-9])?)$`

	// Create a regular expression object
	regex := regexp.MustCompile(validArgumentRegex)

	// Check if the argument matches the pattern
	return regex.MatchString(string(*argument))
}
