package main

func Solve(num uint32) int {
	result := uint32(0)
	mask := uint32(1)
	length := 32

	for index := 0; index < length; index++ {
		result = result + (num >> (length - index - 1) & mask)
	}

	return int(result)
}

func main() {

}
