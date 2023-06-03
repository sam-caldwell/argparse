package argparse

// Preamble - Add another line to the prefix set
func (arg *Arguments) Preamble(line string) *Arguments {
	arg.preambleText = append(arg.preambleText, line)
	return arg
}
