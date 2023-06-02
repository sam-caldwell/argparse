package argparse

import (
	"fmt"
	"testing"
)

func TestListSize(t *testing.T) {
	var list []string
	for i := 1; i <= 10; i++ {
		list = append(list, fmt.Sprintf("%d", i))
		actual := listSize(&list)
		expected := len(list)
		if expected != actual {
			t.Fatalf("%d List size (%d) expected %d", i, actual, expected)
		}
	}
}
