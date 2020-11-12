package main

func Solve(nums []int, target int) []int {
	ret := []int{-1, -1}
	for i := 0; i < len(nums) && nums[i] <= target; i++ {
		if nums[i] == target {
			if ret[0] == -1 {
				ret[0] = i
			}

			ret[1] = i
		}
	}

	return ret
}

func main() {

}
