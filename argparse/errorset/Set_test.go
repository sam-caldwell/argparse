package errorset

import (
	"fmt"
	"testing"
)

func TestSet_Struct(t *testing.T) {
	var s Set
	if len(s.e) != 0 {
		t.Fatal("errorset Set.e must be length 0 by default")
	}
	s.e = append(s.e, fmt.Errorf("test error 1"))
	if len(s.e) != 1 {
		t.Fatal("errorset Set.e size must be 1")
	}
	s.e = append(s.e, fmt.Errorf("test error 2"))
	if len(s.e) != 2 {
		t.Fatal("errorset Set.e size must be 2")
	}
	if s.e[0].Error() != "test error 1" {
		t.Fatal("error 0 mismatch")
	}
	if s.e[1].Error() != "test error 2" {
		t.Fatal("error 1 mismatch")
	}
}
