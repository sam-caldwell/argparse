package parsed

// Count - Return the number of items in the TokenMap
func (arg *TokenMap) Count() int {
	return len(arg.data)
}
