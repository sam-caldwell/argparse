package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed"
	"github.com/sam-caldwell/orderedset"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	descriptors    descriptormap.Map
	programName    string
	preambleText   orderedset.Set
	postscriptText orderedset.Set
	err            orderedset.Set
	results        parsed.TokenMap
}
