package argparse

const (
	shortHelp             = "-h"
	longHelp              = "--help"
	errInvalidArgument    = "invalid argument (%s)"
	errReservedArg        = "argparse %s and %s are reserved"
	errNotInitialized     = "argparse not properly initialized"
	errInvalidArgType     = "invalid argument type %s"
	errTypeMismatch       = "expected %s (%s)"
	errFlagExpectsNoValue = "flag arguments expect no value. %s"
	errMissingArgument    = "missing argument (%s)"
	errInternalError      = "internal error"
	errMissingRequiredArg = "missing required argument (%s)"
	errNotEnoughArgs      = "not enough arguments"
)
