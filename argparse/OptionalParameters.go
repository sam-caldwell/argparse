package argparse

import "github.com/sam-caldwell/go-argparse/v2/types"

// OptionalParameters - Descriptor for an argument
type OptionalParameters struct {
	argType     types.ArgTypes
	argDefault  interface{}
	argHelp     types.ArgHelp
	argRequired bool
}
