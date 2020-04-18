package main

func Solve(num uint32) uint32 {
	var (
		result = uint32(0)
		mask   = uint32(1)
		length = 32
	)

	for index := 0; index < length; index++ {
		pop := (num >> index) & mask
		result = result | (pop << (length - index - 1))
	}

	return result
}

func main() {

}
