package descriptormap

import "github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"

func (m *Map) Get(name string) *descriptor.Descriptor {
	if value, found := m.data[name]; found {
		return &value
	}
	return nil
}
