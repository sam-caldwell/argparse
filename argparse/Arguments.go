package argparse

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	set            []Argument
	programName    string
	preambleText   []string
	postscriptText []string
	err            []error
	value          map[string]interface{}
}
