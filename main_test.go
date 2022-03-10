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

func TestGetWholeStory(t *testing.T) {
	sampleString := "10-hello-20-world"
	expected := "hello world"
	if wholeStory := getWholeStory(sampleString); wholeStory != expected {
		t.Errorf("Expected %s, but got %s", expected, wholeStory)
	}
}

func TestGetStoryStat(t *testing.T) {
	sampleString := "10-hello-20-world-30-hooray-40-yess"
	shortestWord, longestWord, averageLength, avgWords := getStoryStat(sampleString)
	if shortestWord != "yess" {
		t.Errorf("Expected %s, but got %s", "yess", shortestWord)
	}
	if longestWord != "hooray" {
		t.Errorf("Expected %s, but got %s", "hooray", longestWord)
	}
	if averageLength != 5 {
		t.Errorf("Expected %d, but got %d", 5, averageLength)
	}
	for _, word := range avgWords {
		if _, found := FindInArray([]string{"hello", "world"}, word); !found {
			t.Errorf("Expected %s, but not found in the average length words", word)
		}
	}
}

func TestGenerate(t *testing.T) {
	if _, _, err := testValidity(generate(true)); err != nil {
		t.Errorf("Expected valid, but got invalid")
	}
	if _, _, err := testValidity(generate(false)); err == nil {
		t.Errorf("Expected invalid, but got valid")
	}
}
