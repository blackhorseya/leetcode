package main

import (
	"math"
)

func Solve(n int) int {
	if n < 2 {
		return 0
	}

	notPrimes := make([]bool, n)
	max := int(math.Sqrt(float64(n)))
	for flag := 2; flag <= max; flag++ {
		// If flag number is not a prime number then pass
		if notPrimes[flag] {
			continue
		}

		// Here flag number must be prime number
		for multiple := 2 * flag; multiple < n; multiple += flag {
			// Set multiple of flag number is not a prime number
			notPrimes[multiple] = true
		}
	}

	count := 0
	for index, notPrime := range notPrimes {
		// Ignore 0 and 1
		if index <= 1 {
			continue
		}

		// If it is a prime number then add count
		if !notPrime {
			count++
		}
	}

	return count
}

func main() {

}
