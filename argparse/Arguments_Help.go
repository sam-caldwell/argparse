package argparse

// Help - Print Argument usage (help text)
func (arg *Arguments) Help() (text string) {
	//text = fmt.Sprintf("\n%s\n\n", os.Args[0])
	//
	//for _, line := range arg.descriptor {
	//	text += line + "\n"
	//}
	//for argName, argument := range arg.optional {
	//	text += fmt.Sprintf("\t%s [%s] - %s\n",
	//		argName,
	//		argument.argType.String(),
	//		argument.argHelp)
	//}
	return text
}
