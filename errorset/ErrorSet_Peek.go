package errorset

import "fmt"

// Peek - Return the error at position i.
func (err *ErrorSet) Peek(i int) error {
	sz := len((*err).e)
	if (i < 0) || (i > sz) {
		return fmt.Errorf("ErrorSet.Peek() bounds check error %d", i)
	}
	return (*err).e[i]
}
