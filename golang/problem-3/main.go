package main

func Solve(s string) int {
	if len(s) == 1 {
		return 1
	}

	m := make(map[rune]int)
	charset := []rune(s)
	size, ret, slow, fast := len(charset), 0, 0, 0

	for slow < size && fast < size {
		c := charset[fast]
		if _, ok := m[c]; !ok {
			m[c] = fast

			if fast == size-1 {
				return max(ret, fast-slow+1)
			}
			fast++
		} else {
			ret = max(ret, fast-slow)
			m = make(map[rune]int)
			slow++
			fast = slow
		}
	}

	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {

}
