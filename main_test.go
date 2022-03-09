package main

import (
	"testing"
)

type testValidityTest struct {
	arg      string
	expected bool
}

var testValidityTests = []testValidityTest{
	{"12-hello-34-world", true},
	{"12-hello-h4-world", false},
}

func TestTestValidity(t *testing.T) {
	for _, test := range testValidityTests {
		if output := testValidity(test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
