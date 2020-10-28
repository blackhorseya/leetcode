package main

func Solve(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = true
		} else {
			return true
		}
	}

	return false
}

func main() {

}
