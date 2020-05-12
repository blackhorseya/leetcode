package main

func Solve(nums []int) int {
	high := len(nums)
	top := 1
	bottom := len(nums)
	expected := (top + bottom) * high / 2
	actual := 0
	for _, num := range nums {
		actual += num
	}

	return expected - actual
}

func main() {

}
