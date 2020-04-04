package main

func Solve(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var flag = 0
	for pos := 0; pos < len(nums); pos++ {
		if nums[pos] != val {
			nums[flag] = nums[pos]
			flag++
		}
	}

	return flag
}

func main() {
	
}
