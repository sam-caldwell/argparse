package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
)

// GetByPosition - given a position, return the descriptor name and object
func (m *Map) GetByPosition(position int) (*string, *descriptor.Descriptor) {
	for n, a := range m.data {
		fmt.Println(a.GetPosition(), n)
		if a.GetPosition() == position {
			name := n
			argument := a
			fmt.Println(name, argument)
			return &name, &argument
		}
	}
	return nil, nil
}
