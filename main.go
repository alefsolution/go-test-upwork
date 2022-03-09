package main

import (
	"errors"
	"strconv"
	"strings"
)

func main() {
}

func testValidity(arg string) ([]int, []string, error) {
	originalSamples := strings.Split(arg, "-")
	numberSamples := []int{}
	stringSamples := []string{}
	for idx, sample := range originalSamples {
		if idx%2 == 0 {
			numberSample, err := strconv.Atoi(sample)
			if err != nil {
				return nil, nil, errors.New("Given string is not valid")
			}
			numberSamples = append(numberSamples, numberSample)
		} else {
			stringSamples = append(stringSamples, sample)
		}
	}
	return numberSamples, stringSamples, nil
}

func getAverageNumber(arg string) int {
	numberSamples, _, err := testValidity(arg)
	if err != nil {
		panic("Given string is not valid")
	}
	sum := 0
	for i := 0; i < len(numberSamples); i++ {
		sum += numberSamples[i]
	}
	return sum / len(numberSamples)
}

func getWholeStory(arg string) string {
	_, stringSamples, err := testValidity(arg)
	if err != nil {
		panic("Given string is not valid")
	}
	return strings.Join(stringSamples, " ")
}
