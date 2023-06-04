package argument

// GetShort - Search the short arguments and return boolean true if found along with reference to the argument
func (arg *Descriptor) GetShort() string {
	return (*arg)._Long
}
