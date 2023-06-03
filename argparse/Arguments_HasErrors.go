package argparse

// HasErrors - return true if there are any errors in the Arguments struct
func (arg *Arguments) HasErrors() bool {
	return len((*arg).err) > 0
}
