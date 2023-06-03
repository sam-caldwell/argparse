package valid

const (
	validArgRegex      = `^[a-z][a-z0-9_-]*[a-z0-9]$`
	validShortArgRegex = `^-[a-z]$`
	validLongArgRegex  = `^--[a-z][a-z0-9_-]*[a-z0-9]$`
)
