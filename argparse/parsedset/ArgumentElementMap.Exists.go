package parsedset

// Exists - return object if found or nil of not found
func (arg *ArgumentElementMap) Exists(n *string) bool {
	return arg.Lookup(n) != nil
}
