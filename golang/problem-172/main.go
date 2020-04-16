package main

// reference https://brilliant.org/wiki/trailing-number-of-zeros/
func Solve(n int) int {
	result := 0

	for n >= 5 {
		n /= 5
		result += n
	}

	return result
}

func main() {

}
