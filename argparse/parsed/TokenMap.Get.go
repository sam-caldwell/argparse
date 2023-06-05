package parsed

// Get - Return a data value associated with descriptor n.
func (arg *TokenMap) Get(n string) any {
	thisArg := arg.Lookup(&n)
	if thisArg == nil {
		return nil
	}
	return thisArg.GetValue()
}
