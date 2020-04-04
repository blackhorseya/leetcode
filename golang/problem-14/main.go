package main

import (
	"math"
	"strings"
)

func Solve(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var lessLen = math.MaxInt64
	var lessStr = ""
	for _, str := range strs {
		if len(str) < lessLen {
			lessLen = len(str)
			lessStr = str
		}
	}

	for pos := lessLen - 1; pos >= 0; pos-- {
		var subStr = lessStr[:pos + 1]
		var have = true
		for _, v := range strs {
			index := strings.Index(v, subStr)
			if index != 0 {
				have = false
				break
			}
		}
		// fmt.Printf("substr: %s, index is 0: %v\n", subStr, have)

		if have {
			return subStr
		}
	}

	return ""
}

func main() {
	
}
