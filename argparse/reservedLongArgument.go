package argparse

// reservedLongArgument - return true if --help is present
func reservedLongArgument(name *string) bool {
	return *name == argHelpLong
}
