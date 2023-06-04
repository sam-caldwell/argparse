package argparse

import "strings"

// searchOptionalArguments - search the arguments for thisArg and return the Argument record or error
func (arg *Arguments) searchOptionalArguments(thisArg *string) (*Argument, bool) {
	trimmedArg := strings.TrimLeft(*thisArg, hyphen)
	if foundArg, found := arg.GetShort(&trimmedArg); found { //nil error means we found something
		return foundArg, found
	}
	if foundArg, found := arg.GetLong(&trimmedArg); found { //nil error means we found something
		return foundArg, found
	}
	return nil, false
}
