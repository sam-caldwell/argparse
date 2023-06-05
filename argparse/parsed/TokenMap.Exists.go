package parsed

// Exists - return object if found or nil of not found
func (arg *TokenMap) Exists(n *string) bool {
	return arg.Lookup(n) != nil
}
