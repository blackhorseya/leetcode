package main

import (
	"sort"
)

func Solve(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return nil
	}

	var result []int
	var flag = make(map[int]int)
	for pos, num := range nums {
		var need = target - num
		if flag[need] > 0 {
			// found
			result = append(result, pos)
			result = append(result, flag[need] - 1)
			break
		} else {
			// not found
			flag[num] = pos + 1
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	
}
