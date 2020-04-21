package main

func Solve(n int) int {
	if n <= 2 {
		return 0
	}

	// Initialize the counter, at least (n / 2), because half is even
	count := n / 2
	records := make([]bool, n)
	// 3, 5, 7, 9...
	for flag := 3; flag*flag < n; flag += 2 {
		// If records[flag] is not a prime number then pass
		if records[flag] {
			continue
		}

		// Here flag must be prime number
		for index := flag * flag; index < n; index += 2 * flag {
			if !records[index] {
				count--
				records[index] = true
			}
		}
	}

	return count
}

func main() {

}
