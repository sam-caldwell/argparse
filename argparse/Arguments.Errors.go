package argparse

func (arg *Arguments) Errors() (result []string) {
	for _, e := range arg.err.List() {
		result = append(result, e.(error).Error())
	}
	return result
}
