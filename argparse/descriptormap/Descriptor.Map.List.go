package descriptormap

import "github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"

func (m *Map) List() map[string]descriptor.Descriptor {
	if m.data == nil {
		return make(map[string]descriptor.Descriptor)
	}
	return m.data
}
