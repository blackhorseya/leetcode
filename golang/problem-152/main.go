package main

func Solve(nums []int) int {
	ret, minimum, maximum := nums[0], nums[0], nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}

		if num < 0 {
			maximum, minimum = minimum, maximum
		}

		maximum = max(num, maximum*num)
		minimum = min(num, minimum*num)
		ret = max(ret, maximum)
	}

	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {

}
