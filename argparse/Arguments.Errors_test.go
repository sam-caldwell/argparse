package argparse

import (
	"fmt"
	"testing"
)

func TestArguments_Errors(t *testing.T) {
	// Create an instance of the Arguments struct
	var a = Arguments{}
	_ = a.err.Add(fmt.Errorf("error1"))
	_ = a.err.Add(fmt.Errorf("error2"))
	_ = a.err.Add(fmt.Errorf("error3"))
	_ = a.err.Add(fmt.Errorf("error4"))
	_ = a.err.Add(fmt.Errorf("error5"))

	if a.err.Count() != 5 {
		t.Fatal("Setup failed.  expect 5 errors")
	}
	errorList := a.Errors()

	if len(errorList) != 5 {
		t.Fatal("Expected count 5 not recd")
	}
	if errorList[1] != "error1" {
		t.Fatal("error1 not found")
	}
	if errorList[2] != "error2" {
		t.Fatal("error2 not found")
	}
	if errorList[3] != "error3" {
		t.Fatal("error3 not found")
	}
	if errorList[4] != "error4" {
		t.Fatal("error4 not found")
	}
	if errorList[5] != "error5" {
		t.Fatal("error4 not found")
	}
}
