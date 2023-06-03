package argparse

const (
	// Exit codes
	exitArgParseError = 1

	// Reserved arguments
	argHelpShort = "-h"
	argHelpLong  = "--help"
	argHelpName  = "help"
	hyphen       = "-"

	// Error messages
	// -- Argument struct methods
	errMissingDefault   = "missing default value (type: %s)"
	errEmptyOrNilObject = "missing or nil input"
	errParsingCliArgs   = "Error parsing command-line arguments"
	// --
	errInvalidArgument    = "invalid argument (%s)"
	errReservedArg        = "commandline argument %s is reserved"
	errNotInitialized     = "argparse not properly initialized"
	errTypeMismatch       = "expected %s (%s)"
	errMissingArgument    = "missing argument (%s)"
	errInternalError      = "internal error"
	errMissingRequiredArg = "missing required argument (%s)"
	errNotEnoughArgs      = "not enough arguments"
	errPositionalArgFirst = "positional arguments must be before any optional arguments"
)
