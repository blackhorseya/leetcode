package main

func Solve(s string) int {	var length = 0
	var found = false
	for pos := len(s) - 1; pos >= 0; pos-- {
		if s[pos] == ' ' {
			if found {
				break
			}

			continue
		}

		length++
		found = true
	}

	return length
}

func main() {
	
}
