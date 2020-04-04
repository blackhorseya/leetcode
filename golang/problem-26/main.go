package main

func Solve(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	var i = 0
	for pos := 1; pos < len(nums); pos++ {
		if nums[pos] != nums[i] {
			i++
			nums[i] = nums[pos]
		}
	}

	return i + 1
}

func main() {
	
}
