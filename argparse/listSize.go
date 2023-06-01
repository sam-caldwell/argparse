package argparse

// listSize - return the size of a string list
func listSize(list *[]string) int {
	if list == nil {
		return 0
	}
	return len(*list)
}
