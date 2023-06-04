package errorset

import "fmt"

// Push - Add an error to ErrorSet array
func (err *ErrorSet) Push(m string) {
	(*err).e = append((*err).e, fmt.Errorf(m))
}
