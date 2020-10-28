package main

func Solve(s string) int {
	// ascii code
	asc := make([]int, 128)
	left := 0
	ret := 0

	for i, c := range s {
		if asc[c] > left {
			left = asc[c]
		}

		if n := i - left + 1; n > ret {
			ret = n
		}

		asc[c] = i + 1
	}

	return ret
}

func main() {

}
