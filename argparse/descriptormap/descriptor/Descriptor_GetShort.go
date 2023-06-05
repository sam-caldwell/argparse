package descriptor

// GetShort - Search the short arguments and return boolean true if found along with reference to the descriptor
func (arg *descriptor.Descriptor) GetShort() string {
	return arg.short
}
