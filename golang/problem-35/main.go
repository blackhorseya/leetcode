package main

func Solve(nums []int, target int) int {
	var low = 0
	var high = len(nums)

	for low < high {
		var mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func main() {
	
}
