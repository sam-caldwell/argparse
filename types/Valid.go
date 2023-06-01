package types

// Valid - Return true if ArgTypes is valid or false if ArgTypes is not valid.
func (arg *ArgTypes) Valid() (result bool) {
	switch *arg {
	case String:
		fallthrough
	case Integer:
		fallthrough
	case Boolean:
		fallthrough
	case Float:
		fallthrough
	case Flag:
		result = true
	default:
		result = false
	}
	return result
}
