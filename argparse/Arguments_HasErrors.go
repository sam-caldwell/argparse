package argparse

func (arg *Arguments) HasErrors() bool {
	return len((*arg).err) > 0
}
