package argparse

// ListErrors - List any errors and return the string list
func (arg *Arguments) ListErrors() (result []string) {
	return arg.err.List()
}
