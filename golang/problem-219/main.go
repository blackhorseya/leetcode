package main

func Solve(nums []int, k int) bool {
	m := make(map[int]int)

	for index, num := range nums {
		if pos, ok := m[num]; !ok {
			m[num] = index
		} else {
			if index-pos <= k {
				return true
			}

			m[num] = index
		}
	}

	return false
}

func main() {

}
