package argparse

// Postscript - Add another line to the postfix set
func (arg *Arguments) Postscript(line string) *Arguments {
	if err := arg.postscriptText.Add(line); err != nil {
		arg.err.Push(err)
	}
	return arg
}
