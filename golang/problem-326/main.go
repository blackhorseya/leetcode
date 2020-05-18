package main

const base = 3

func Solve(n int) bool {
	for n > 1 {
		if n % base != 0 {
			return false
		}

		n /= base
	}

	if n == 1 {
		return true
	}

	return false
}

func main() {

}
