package errorset

import (
	"fmt"
	"testing"
)

func TestSet_Pop(t *testing.T) {
	var s Set
	if s.Count() != 0 {
		t.Fatal("errorset Set.e must be length 0 by default")
	}

	if err := s.Peek(0); err != nil {
		t.Fatal(err)
	}
	if err := s.Pop(); err != nil {
		t.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		actualError := fmt.Sprintf("test error %d", i)
		s.e = append(s.e, fmt.Errorf(actualError))
		if s.Count() != i {
			t.Fatalf("error %d: %s", i, actualError)
		}
		if err := s.Peek(i - 1); err.Error() != actualError {
			t.Fatalf("error mismatch (%d): %s", i, err)
		}
	}

	for i := 1; i <= 10; i++ {
		if err := s.Pop(); err.Error() != fmt.Sprintf("test error %d", i) {
			t.Fatalf("pop error mismatch: %s", err.Error())
		}
	}
}
