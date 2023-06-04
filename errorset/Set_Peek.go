package errorset

// Peek - Return the error at position i or a bounds check error
func (err *Set) Peek(i int) error {
	if (i < 0) || (i > err.Count()) {
		panic(errBoundsCheckError)
	}
	return (*err).e[i]
}
