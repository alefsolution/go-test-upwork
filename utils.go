package main

import (
	"math/rand"
	"strings"
	"time"
)

func FindInArray(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

type SampleType int

const (
	ALPHABETIC = iota
	NUMERIC
)

func RandString(sampleType SampleType) string {
	var charSet []rune
	switch sampleType {
	case NUMERIC:
		charSet = []rune("1234567890")
	case ALPHABETIC:
		charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}

	var output strings.Builder
	s := rand.NewSource(time.Now().UnixNano())
	length := rand.New(s).Intn(10) + 2

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}

	return output.String()
}
