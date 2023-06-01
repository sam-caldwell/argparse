package argparse

import "testing"

func TestIsPositional(t *testing.T) {
	args := map[string]bool{
		"arg":   true,
		"arg-":  true,
		"arg--": true,
		"a-rg":  true,
		"ar--g": true,
		"--arg": false,
		"-arg":  true,
		"-a":    false,
	}
	for thisArg, expectedResult := range args {
		actualResult := isPositional(&thisArg)
		if expectedResult {
			if actualResult {
				t.Logf("OK  :1: arg (%s) expected: %v actual: %v", thisArg, expectedResult, actualResult)
			} else {
				t.Fatalf("FAIL:1: arg (%s) expected: %v actual: %v", thisArg, expectedResult, actualResult)
			}
		} else {
			if !actualResult {
				t.Logf("OK  :2: arg (%s) expected: %v actual: %v", thisArg, expectedResult, actualResult)
			} else {
				t.Fatalf("FAIL:2: arg (%s) expected: %v actual: %v", thisArg, expectedResult, actualResult)
			}
		}
	}
}
