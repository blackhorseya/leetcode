package main

func Solve(nums []int) []int {
	var ret []int
	size := len(nums)
	counters := make([]int, size + 1)

	for _, num := range nums {
		counters[num] += 1
	}

	for index, count := range counters {
		if count == 0 && index != 0 {
			ret = append(ret, index)
		}
	}

	return ret
}

func main() {

}
