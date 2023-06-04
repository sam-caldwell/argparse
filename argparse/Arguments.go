package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/argument"
	"github.com/sam-caldwell/argparse/v2/argparse/parsed"
	"github.com/sam-caldwell/argparse/v2/errorset"
)

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	descriptors    argument.DescriptorMap
	programName    string
	preambleText   []string
	postscriptText []string
	err            errorset.Set
	results        parsed.Arguments
}
