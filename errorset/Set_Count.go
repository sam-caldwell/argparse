package errorset

// Count - return a count of the elements in the set
func (err *Set) Count() int {
	return len((*err).e)
}
