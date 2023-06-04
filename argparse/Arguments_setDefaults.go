package argparse

func (arg *Arguments) setDefaults() (err error) {
	for _, argument := range arg.positional {
		arg.value[argument.argName] = argument.argDefault
	}
	for name, argument := range arg.optional {
		arg.value[name] = argument.argDefault
	}
	return nil
}
