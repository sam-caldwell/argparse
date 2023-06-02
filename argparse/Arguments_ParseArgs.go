package argparse

// ParseArgs - Parse os.Args and store the state internally
func (args *Arguments) ParseArgs() (err error) {
	args.Debug("ParseArgs() started")
	if err = args.setDefaults(); err != nil {
		return err
	}
	args.Debug("ParseArgs(): setDefaults() ok")

	if err = args.parsePositionalArgs(); err != nil {
		return err
	}
	args.Debug("ParseArgs(): parsePositionalArgs() ok")

	if err = args.parseOptionalArgs(); err != nil {
		return err
	}
	args.Debug("ParseArgs(): parseOptionalArgs() ok")

	defer args.Debugf("ParseArgs() done.  error:%v", err)
	return err
}
