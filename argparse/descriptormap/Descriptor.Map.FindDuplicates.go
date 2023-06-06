package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/simpleset/v2"
)

func (m *Map) FindDuplicates(d descriptor.Descriptor) error {
	var argSet simpleset.Set
	var shortSet simpleset.Set
	var longSet simpleset.Set

	for name, argument := range m.data {
		if err := argSet.Add(name); err != nil {
			return fmt.Errorf("argument cannot be redefined %s", name)
		}
		if err := shortSet.Add(argument.GetShort()); err != nil {
			return fmt.Errorf("short argument cannot be redefined %s", name)
		}
		if err := longSet.Add(argument.GetLong()); err != nil {
			return fmt.Errorf("long argument cannot be redefined %s", name)
		}
	}
	return nil
}
