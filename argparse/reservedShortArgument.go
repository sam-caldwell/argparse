package argparse

// reservedShortArgument - return true if -h is present
func reservedShortArgument(name *string) bool {
	return *name == argHelpShort
}
