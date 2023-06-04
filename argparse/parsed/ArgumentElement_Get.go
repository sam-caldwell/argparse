package parsed

// Get - Return the contents of an ArgumentElement object.
func (arg *ArgumentElement) Get() (string, any) {
	return arg._name, arg._value
}
