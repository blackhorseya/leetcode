package main

import (
	"strconv"
)

func Solve(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	var result = ""
	var str = Solve(n - 1)
	var count = 1
	for pos := 1; pos < len(str); pos++ {
		if str[pos] == str[pos-1] {
			count += 1
		} else {
			result += strconv.Itoa(count) + string(str[pos-1])
			count = 1
		}
	}

	result += strconv.Itoa(count) + string(str[len(str) - 1])
	return result
}

func main() {
	
}
