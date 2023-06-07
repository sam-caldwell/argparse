package descriptormap

import "github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"

func (m *Map) GetByName(token *string) (*string, *descriptor.Descriptor) {
	return token, m.Get(token)
}
