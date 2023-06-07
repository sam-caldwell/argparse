package descriptormap

func (m *Map) CountPositionalArgs() (count int) {
	for _, argument := range m.data {
		if (argument.GetShort() == "") || (argument.GetLong() == "") {
			count++
		}
	}
	return count
}
