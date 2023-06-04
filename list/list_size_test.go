package list

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/list"
	"testing"
)

func TestListSize(t *testing.T) {
	var data []string
	for i := 1; i <= 10; i++ {
		data = append(data, fmt.Sprintf("%d", i))
		actual := list.Sizep(&data)
		expected := len(data)
		if expected != actual {
			t.Fatalf("%d List size (%d) expected %d", i, actual, expected)
		}
	}
}
