package argparse

import "fmt"

// ListErrors - List any errors and return the string list
func (arg *Arguments) ListErrors() (result []string) {
	if arg.hasErrors() {

		fmt.Println("Error parsing command-line arguments")
		for _, err := range arg.err {
			if err != nil {
				result = append(result, err.Error())
			}
		}
	}
	return result
}