package argparse

// Postscript - Add another line to the postfix set
func (arg *Arguments) Postscript(line string) *Arguments {
	arg.postscriptText = append(arg.postscriptText, line)
	return arg
}
