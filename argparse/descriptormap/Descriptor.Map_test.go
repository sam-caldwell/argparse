package descriptormap

import (
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"testing"
)

func TestDescriptorMap(t *testing.T) {

	initialTest := func() *Map {
		var d Map
		if d != nil {
			t.Fatal("initial state should be nil")
		}
		d = make(Map)
		if d == nil {
			t.Fatal("descriptor map should be initialized")
		}
		return &d
	}

	test := func(d *Map, s string, sz int) {
		if *d == nil {
			*d = make(Map)
		}
		(*d)[s] = descriptor.Descriptor{}
		if len(*d) != sz {
			t.Fatalf("expected size %d not realized (%d): %s", len(*d), sz, s)
		}
	}

	d := initialTest()
	test(d, "test1", 1)
	test(d, "test2", 2)
	test(d, "test3", 3)
	test(d, "test4", 4)
	test(d, "test5", 5)
}
