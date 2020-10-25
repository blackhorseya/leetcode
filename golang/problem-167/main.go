package main

func Solve(numbers []int, target int) []int {
	ret := make([]int, 2)
	m := make(map[int]int)

	for first, number := range numbers {
		if second, ok := m[target-number]; ok && second != 0 {
			ret[0] = second
			ret[1] = first + 1

			break
		}

		m[number] = first + 1
	}

	return ret
}

func main() {

}
