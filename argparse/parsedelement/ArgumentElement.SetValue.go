package parsedelement

// SetValue  - Set the given name Argument value for an ArgumentElement
func (arg *ArgumentElement) SetValue(v any) (err error) {
	//We need a typecheck before setting the value
	if err := arg.typ.Typecheck(v); err != nil {
		return err
	}
	arg.value = v
	return nil
}
