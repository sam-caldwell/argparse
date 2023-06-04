package valid

const (
	validArgRegex      = `^[a-z][a-z0-9_-]*[a-z0-9]$`
	validShortArgRegex = `^-[a-z]$`
	validLongArgRegex  = `^--[a-z][a-z0-9_-]*[a-z0-9]$`
	validYearRegex     = `^[0-9]{4}$`
)

const (
	errEmptyOrNilObject = "missing or nil input"
	errExpectedShortArg = "expected short argument for %s"
	errExpectedLongArg  = "expected long argument for %s"
	errExpectedNameArg  = "expected name argument for %s"

	errInvalidYear = "invalid year %d"
)
