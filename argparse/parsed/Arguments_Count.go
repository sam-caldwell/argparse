package parsed

// Count - Count the number of records in the parsed.Arguments
func (arg *Arguments) Count() int {
	if arg.set != nil {
		return len(arg.set)
	}
	return 0
}
