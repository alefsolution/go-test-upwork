package main

import (
	"testing"
)

var testValidityTests = []string{
	"12-hello-34-world",
	"12-hello-h4-world",
}

func TestTestValidity(t *testing.T) {
	if _, _, err := testValidity(testValidityTests[0]); err != nil {
		t.Errorf("Validation should success but failed")
	}
	if _, _, err := testValidity(testValidityTests[1]); err == nil {
		t.Errorf("Validation should fail but succeeded")
	}
}

func TestGetAverageNumber(t *testing.T) {
	sampleString := "10-hello-20-world"
	expected := 15
	if avg := getAverageNumber(sampleString); avg != expected {
		t.Errorf("Expected %d, but got %d", expected, avg)
	}
}
