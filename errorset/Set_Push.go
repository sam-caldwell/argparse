package errorset

import "fmt"

// Push - Add an error to Set array
func (err *Set) Push(m string) {
	(*err).e = append((*err).e, fmt.Errorf(m))
}
