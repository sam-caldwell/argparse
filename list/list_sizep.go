package list

// Sizep - Return the size of the given list (by reference)
func Sizep(list *[]any) int {
	if list == nil {
		return 0
	}
	return len(*list)
}
