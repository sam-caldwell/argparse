package argparse

// GetLong - Search the long arguments and return boolean true if found along with reference to the argument
func (arg *Arguments) GetLong(argument *string) (*Argument, bool) {
	for _, thisArgument := range arg.set {
		if *argument == thisArgument._Short {
			return &thisArgument, true
		}
	}
	return nil, false

}
