package parsedset

// Get - Return a data value associated with argument n.
func (arg *ArgumentElementMap) Get(n string) any {
	thisArg := arg.Lookup(&n)
	if thisArg == nil {
		return nil
	}
	return thisArg.GetValue()
}
