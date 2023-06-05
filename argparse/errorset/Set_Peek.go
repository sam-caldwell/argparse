package errorset

import "fmt"

// Peek - Return the error at position i or a bounds check error
func (err *Set) Peek(i int) error {
	if (i < 0) || (i > err.Count()) {
		return fmt.Errorf(errBoundsCheckError)
	}
	if err.Count() == 0 {
		return nil
	}
	return (*err).e[i]
}
