package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap"
	"github.com/sam-caldwell/argparse/v2/argparse/errorset"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	descriptors    descriptormap.Map
	programName    string
	preambleText   []string
	postscriptText []string
	err            errorset.Set
	results        parsed.TokenMap
}
