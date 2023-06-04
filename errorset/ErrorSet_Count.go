package errorset

// Count - return a count of the elements in the set
func (err *ErrorSet) Count() int {
	return len((*err).e)
}
