package main

func Solve(nums []int, target int) []int {
	return []int{find(nums, target, -1), find(nums, target, 1)}
}

func find(nums []int, target int, direction int) int {
	low, high, ret := 0, len(nums)-1, -1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			ret = mid

			if direction == 1 {
				low = mid + 1
			} else if direction == -1 {
				high = mid - 1
			}
		}
	}

	return ret
}

func main() {

}
