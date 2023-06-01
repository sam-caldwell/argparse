package argparse

// ParseArgs - Parse os.Args and store the state internally
func (args *Arguments) ParseArgs() (err error) {

	if err = args.setDefaults(); err != nil {
		return err
	}

	if err = args.parsePositionalArgs(); err != nil {
		return err
	}

	if err = args.parseOptionalArgs(); err != nil {
		return err
	}

	return err
}
