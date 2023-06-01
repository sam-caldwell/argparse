package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// searchStringList - linear search of a string list
func searchStringList(expected *types.ArgString, source *[]string) (string, bool) {
	for _, actual := range *source {
		if actual == string(*expected) {
			return actual, true
		}
	}
	return "", false
}
