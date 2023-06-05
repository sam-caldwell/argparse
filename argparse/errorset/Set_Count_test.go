package errorset

import (
	"fmt"
	"testing"
)

func TestSet_Count(t *testing.T) {
	var s Set
	if s.Count() != 0 {
		t.Fatal("errorset Set.e must be length 0 by default")
	}
	s.e = append(s.e, fmt.Errorf("test error 1"))
	if s.Count() != 1 {
		t.Fatal("errorset Set.e size must be 1")
	}
	s.e = append(s.e, fmt.Errorf("test error 2"))
	if s.Count() != 2 {
		t.Fatal("errorset Set.e size must be 2")
	}
	if s.e[0].Error() != "test error 1" {
		t.Fatal("error 0 mismatch")
	}
	if s.e[1].Error() != "test error 2" {
		t.Fatal("error 1 mismatch")
	}
}
