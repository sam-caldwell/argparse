package argparse

import "github.com/sam-caldwell/go-argparse/v2/errorset"

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	set            []Argument
	programName    string
	preambleText   []string
	postscriptText []string
	err            errorset.Set
	value          map[string]interface{}
}
