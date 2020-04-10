package main

import (
	"strings"
)

func isAlphanumeric(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func Solve(s string) bool {
	s = strings.ToLower(s)

	left, right := 0, len(s) - 1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {

}
