package main

func Solve(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	nums[1] = max(nums[0], nums[1])
	for flag := 2; flag < length; flag++ {
		slow := flag - 2
		fast := flag - 1
		nums[flag] = max(nums[flag]+nums[slow], nums[fast])
	}
	return nums[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

}
