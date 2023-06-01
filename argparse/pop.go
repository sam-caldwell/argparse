package argparse

import "fmt"

// pop - pop lowest element from the given list (string array)
func pop(list *[]string) (value string, err error) {
	if listSize(list) > 0 {
		defer func() {
			*list = (*list)[1:]
		}()
		return (*list)[0], nil
	}
	return "", fmt.Errorf(errNotEnoughArgs)
}
