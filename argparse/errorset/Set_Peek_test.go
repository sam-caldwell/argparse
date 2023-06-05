package errorset

import (
	"fmt"
	"testing"
)

func TestSet_Peek(t *testing.T) {
	var s Set
	if s.Count() != 0 {
		t.Fatal("errorset Set.e must be length 0 by default")
	}

	if err := s.Peek(0); err != nil {
		t.Fatal(err) // Expect nil error on empty set
	}

	if err := s.Peek(-1); err.Error() != errBoundsCheckError {
		t.Fatal(err) //Expected bounds-check error
	}

	actualError1 := "test error 1"
	s.e = append(s.e, fmt.Errorf(actualError1))
	if s.Count() != 1 {
		t.Fatal("errorset Set.e size must be 1")
	}
	if err := s.Peek(0); err.Error() != actualError1 {
		t.Fatalf("errorset Set.e mismatched to expected error:%s", actualError1)
	}

	actualError2 := "test error 2"
	s.e = append(s.e, fmt.Errorf(actualError2))
	if s.Count() != 2 {
		t.Fatal("errorset Set.e size must be 2")
	}
	if err := s.Peek(0); err.Error() != actualError1 {
		t.Fatalf("errorset Set.e mismatched to expected error:%s", actualError1)
	}
	if err := s.Peek(1); err.Error() != actualError2 {
		t.Fatalf("errorset Set.e mismatched to expected error:%s", actualError2)
	}

	if s.Peek(0).Error() != "test error 1" {
		t.Fatal("error 0 mismatch")
	}
	if s.Peek(1).Error() != "test error 2" {
		t.Fatal("error 1 mismatch")
	}
}
