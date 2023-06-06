package descriptor

// GetType - Return current type
func (arg *Descriptor) GetType() string {
	return arg.typ.String()
}
