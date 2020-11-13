package main

import (
	"math"
)

var toInt = map[byte]int64{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func Solve(s string) int {
	ret, factor := int64(0), int64(1)
	i, l := 0, len(s)

	// judgment space character
	for i < l && s[i] == ' ' {
		i++
	}

	// judgment sign symbol
	if i < l && s[i] == '-' {
		factor = -1
		i++
	} else if i < l && s[i] == '+' {
		i++
	}

	// judgment digit character
	for i < l && ret < math.MaxInt32 {
		if val, ok := toInt[s[i]]; ok {
			ret = ret * 10 + val
			i++
		} else {
			break
		}
	}

	// added sign
	ret = ret * factor

	// judgment maximum
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}

	// judgment minimum
	if ret < math.MinInt32 {
		return math.MinInt32
	}

	return int(ret)
}

func main() {

}
