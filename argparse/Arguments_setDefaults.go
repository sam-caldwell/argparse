package argparse

func (args *Arguments) setDefaults() (err error) {
	for _, argument := range args.positional {
		args.value[argument.argName] = argument.argDefault
	}
	for name, argument := range args.optional {
		args.value[name] = argument.argDefault
	}
	return nil
}
