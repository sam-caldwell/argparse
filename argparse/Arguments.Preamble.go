package argparse

// Preamble - Add another line to the prefix set
func (arg *Arguments) Preamble(line string) *Arguments {
	if err := arg.preambleText.Add(line); err != nil {
		_ = arg.err.Push(err)
	}
	return arg
}
