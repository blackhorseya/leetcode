package main

func Solve(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := nums[0]
	for slow := 0; slow < len(nums); slow++ {
		temp := 1
		for fast := slow; fast < len(nums); fast++ {
			temp *= nums[fast]
			if temp > ret {
				ret = temp
			}
		}
	}

	return ret
}

func main() {

}
