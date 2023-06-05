package descriptormap

func (m *Map) Count() int {
	if m.data == nil {
		return 0
	}
	return len(m.data)
}
