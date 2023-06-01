package argparse

import (
	"strings"
)

// isPositional - return boolean if argument is positional (it has no - or -- starting chars).
func isPositional(argument *string) bool {
	const (
		shortStart = "-"
		longStart  = "--"
	)
	if argument == nil {
		return false
	}
	trimmed := strings.TrimSpace(*argument)

	if strings.HasPrefix(trimmed, longStart) && (len(trimmed) > 4) {
		return false
	}
	if strings.HasPrefix(trimmed, shortStart) && (len(trimmed) == 2) {
		return false
	}
	return true
}
