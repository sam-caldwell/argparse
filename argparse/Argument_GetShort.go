package argparse

// GetShort - Search the short arguments and return boolean true if found along with reference to the argument
func (arg *Arguments) GetShort(argument *string) (*Argument, bool) {
	for _, thisArgument := range arg.set {
		if *argument == thisArgument._Short {
			return &thisArgument, true
		}
	}
	return nil, false

}
