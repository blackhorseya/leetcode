package main

func Solve(nums []int) int {
	maps := make(map[int]bool)
	for _, num := range nums {
		maps[num] = true
	}
	for index := 0; index < (len(nums) + 1); index++ {
		_, ok := maps[index]
		if !ok {
			return index
		}
	}

	return 0
}

func main() {

}
