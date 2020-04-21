package main

func Solve(n int) bool {
	sum := 0
	mod := 0
	for ; n != 0; n /= 10 {
		mod = n % 10
		sum += mod * mod
	}

	if sum < 10 && sum == mod*mod {
		return sum == 1
	} else {
		return Solve(sum)
	}
}

func main() {

}
