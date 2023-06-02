package argparse

const (
	//Reserved arguments
	argHelpShort = "-h"
	argHelpLong  = "--help"
	argHelp      = "help"

	// Error messages
	errInvalidArgument    = "invalid argument (%s)"
	errReservedArg        = "argparse %s and %s are reserved"
	errNotInitialized     = "argparse not properly initialized"
	errInvalidArgType     = "invalid argument type %s"
	errTypeMismatch       = "expected %s (%s)"
	errMissingArgument    = "missing argument (%s)"
	errInternalError      = "internal error"
	errMissingRequiredArg = "missing required argument (%s)"
	errNotEnoughArgs      = "not enough arguments"
)
