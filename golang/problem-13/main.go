package main

func Solve(s string) int {
	if len(s) == 0 {
		return 0
	}

	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result = roman[string(s[len(s) - 1])]
	for pos := len(s) - 2; pos >= 0; pos-- {
		var prev = roman[string(s[pos + 1])]
		var cur = roman[string(s[pos])]
		if prev > cur {
			result -= cur
		} else {
			result += cur
		}
	}
	return result
}

func main() {
	
}
