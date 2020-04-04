package main

import (
	"math"
)

func Solve(x int) int {
	// 只有單一個數字
	if x >= 0 && x <= 9 {
		return x
	}

	var result = 0
	for index := x; index != 0; index = index/10 {
		result = result*10 + index%10
	}

	// 超過 int32 大小
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}

func main() {
	
}
