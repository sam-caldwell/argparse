package argument

// GetLong - Search the long arguments and return boolean true if found along with reference to the argument
func (arg *Descriptor) GetLong() string {
	return (*arg).short
}
