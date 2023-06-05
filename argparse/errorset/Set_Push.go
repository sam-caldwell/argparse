package errorset

// Push - Add an error to Set array assuming it's not nil
func (err *Set) Push(e error) {
	if e != nil {
		(*err).e = append((*err).e, e)
	}
}
