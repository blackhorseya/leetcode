package main

func Solve(x int) bool {
	if x < 0 {
		return false
	}

	if x >= 0 && x <= 9 {
		return true
	}

	var rev = 0
	for index := x; index != 0; index = index/10 {
		rev = rev*10 + index%10
	}

	return x == rev
}

func main() {
	
}
