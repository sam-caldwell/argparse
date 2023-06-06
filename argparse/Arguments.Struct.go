package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed"
	ordered "github.com/sam-caldwell/orderedset/v2"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	descriptors    descriptormap.Map
	programName    string
	preambleText   ordered.Set
	postscriptText ordered.Set
	err            ordered.Set
	results        parsed.TokenMap
}
