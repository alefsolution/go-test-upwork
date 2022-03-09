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

func getStoryStat(arg string) (string, string, int, []string) {
	_, stringSamples, err := testValidity(arg)
	if err != nil {
		panic("Given string is not valid")
	}

	shortestWord := stringSamples[0]
	longestWord := stringSamples[0]
	totalLength := 0

	for _, word := range stringSamples {
		totalLength += len(word)
		if len(word) < len(shortestWord) {
			shortestWord = word
		}
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}
	averageLength := totalLength / len(stringSamples)

	var avgWords []string
	averageLength = totalLength / len(stringSamples)
	for _, word := range stringSamples {
		if len(word) == averageLength {
			avgWords = append(avgWords, word)
		}
	}

	return shortestWord, longestWord, averageLength, avgWords
}
