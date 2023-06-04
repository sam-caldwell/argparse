package errorset

import "fmt"

// Pushf - Add an error to Set array assuming it's not nil using a format string
func (err *Set) Pushf(format string, e any) {
	if e != nil {
		(*err).e = append((*err).e, fmt.Errorf(format, e))
	}
}
