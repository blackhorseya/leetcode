package main

import (
	"math"
	"strconv"
	"strings"
)

func Solve(s string) int {
	s = strings.TrimLeft(s, " ")

	if len(s) == 0 {
		return 0
	}

	if !IsSign(rune(s[0])) && !IsDigit(rune(s[0])) {
		return 0
	}

	found := false
	var msg []rune
	for i, c := range s {
		if IsSign(c) || IsDigit(c) {
			found = true
			msg = append(msg, c)
		}

		if found && i+1 < len(s) && !IsDigit(rune(s[i+1])) {
			break
		}
	}

	ret, _ := strconv.Atoi(string(msg))
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}

	if ret < math.MinInt32 {
		return math.MinInt32
	}

	return ret
}

func IsDigit(c rune) bool {
	if c > '9' || c < '0' {
		return false
	}

	return true
}

func IsSign(c rune) bool {
	if c != '+' && c != '-' {
		return false
	}

	return true
}

func main() {

}
