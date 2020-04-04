package main

func Solve(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max = nums[0]
	var sum = max
	for _, v := range nums[1:] {
		if sum + v > v {
			sum = sum + v
		} else {
			sum = v
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

func main() {
	
}
