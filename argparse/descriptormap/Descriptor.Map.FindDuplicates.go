package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/simpleset/v2"
)

func (m *Map) FindDuplicates() error {
	var argSet simpleset.Set
	var shortSet simpleset.Set
	var longSet simpleset.Set

	for name, argument := range m.data {
		if err := argSet.Add(name); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
		if err := shortSet.Add(argument.GetShort()); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
		if err := longSet.Add(argument.GetLong()); err != nil {
			return fmt.Errorf(errArgumentCannotBeRedefined, name)
		}
	}
	return nil
}
