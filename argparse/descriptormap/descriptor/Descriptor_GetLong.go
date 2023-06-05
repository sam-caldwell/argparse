package descriptor

// GetLong - Search the long arguments and return boolean true if found along with reference to the descriptor
func (arg *descriptor.Descriptor) GetLong() string {
	return arg.long
}
