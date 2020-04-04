package main

import (
	"strings"
)

func Solve(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	return strings.Index(haystack, needle)
}

func main() {
	
}
