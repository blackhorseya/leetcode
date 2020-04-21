package main

import (
	"math"
)

var visited= make(map[int]int)

func Solve(n int) bool {
	sum := 0
	for (n / 10) > 0 {
		sum = sum + int(math.Pow(float64(n%10), 2))
		n /= 10
	}
	sum = sum + int(math.Pow(float64(n%10), 2))

	if sum == 1 {
		return true
	}

	if visited[sum] == 0 {
		visited[sum] = 3
		if Solve(sum) {
			visited[sum] = 1
			return true
		} else {
			visited[sum] = 2
			return false
		}
	} else {
		if visited[sum] == 1 {
			return true
		} else {
			return false
		}
	}
}

func main() {

}
