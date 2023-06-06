package argparse

func (arg *Arguments) ErrorCount() int {
	return arg.err.Count()
}
