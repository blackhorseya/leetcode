package main

func Solve(numbers []int, target int) []int {
	var ret []int
	head, tail := 0, len(numbers)-1

	for head < tail {
		sum := numbers[head] + numbers[tail]

		if sum == target {
			return []int{head + 1, tail + 1}
		} else if sum > target {
			tail--
		} else if sum < target {
			head++
		}
	}

	return ret
}

func main() {

}
