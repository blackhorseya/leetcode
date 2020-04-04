package main

func Solve(digits []int) []int {
	if digits == nil || len(digits) == 0 {
		return []int{1}
	}


	for pos := len(digits) - 1; pos >= 0; pos-- {
		digits[pos] = digits[pos] + 1
		if digits[pos] / 10 == 0 {
			return digits
		} else {
			digits[pos] = 0
		}
	}

	return append([]int{1}, digits...)
}

func main() {
	
}
