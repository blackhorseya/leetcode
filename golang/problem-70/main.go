package main

func Solve(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	var s1, s2 = 1, 2
	var r1 = s2

	for pos := 3; pos <= n; pos++ {
		r1 = s1 + s2
		s1 = s2
		s2 = r1
	}

	return r1
}

func main() {
	
}
