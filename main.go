package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testValidity("12-hello-35-world")
}

func testValidity(arg string) bool {
	originalSamples := strings.Split(arg, "-")
	for idx, sample := range originalSamples {
		if idx%2 == 0 {
			if _, err := strconv.Atoi(sample); err != nil {
				fmt.Println("it's not valid")
				return false
			}
		}
	}
	return true
}
