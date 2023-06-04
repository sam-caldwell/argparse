package errorset

// Pop - Return ErrorSet[0] and remove this element
func (err *ErrorSet) Pop() error {
	if len((*err).e) > 0 {
		defer func() {
			(*err).e = append((*err).e[:0], (*err).e[1:]...)
		}()
		return (*err).e[0]
	}
	return nil
}
