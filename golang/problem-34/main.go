package main

func Solve(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	ret := []int{-1, -1}
	for i, num := range nums {
		if num == target {
			if ret[0] == -1 {
				ret[0] = i
			}

			if ret[1] < i {
				ret[1] = i
			}
		}

		if num != target && ret[0] != -1 {
			break
		}
	}

	return ret
}

func main() {

}
