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

func getAverageNumber(arg string) {
}
