package argparse

import "testing"

func TestPop(t *testing.T) {
	testList := []string{"a", "b", "c"}
	if value, err := pop(&testList); err != nil {
		t.Fatal("pop failed (unexpected)")
	} else {
		if value != "a" {
			t.Fatalf("expected a, got %s", value)
		}
	}
	if value, err := pop(&testList); err != nil {
		t.Fatal("pop failed (unexpected)")
	} else {
		if value != "b" {
			t.Fatalf("expected b, got %s", value)
		}
	}
	if value, err := pop(&testList); err != nil {
		t.Fatal("pop failed (unexpected)")
	} else {
		if value != "c" {
			t.Fatalf("expected c, got %s", value)
		}
	}
	if value, err := pop(&testList); err != nil {
		if value != "" {
			t.Fatal("Expected '' on empty list with error")
		}
		if err.Error() != errNotEnoughArgs {
			t.Fatalf("Expected %s but got %s", errNotEnoughArgs, err)
		}
	} else {
		t.Fatalf("expected error on empty")
	}

}
