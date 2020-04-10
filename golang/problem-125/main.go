package main

import (
	"regexp"
	"strings"
)

func Solve(s string) bool {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "")

	length := len(s)
	result := true
	if length == 0 {
		return true
	}

	for index, str := range s[:length/2] {
		if !result {
			break
		}

		revIndex := length - 1 - index
		if string(str) != string(s[revIndex]) {
			result = false
		}
	}

	return result
}

func main() {

}
