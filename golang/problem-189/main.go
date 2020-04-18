package main

func Solve(nums []int, k int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	var tmp []int
	tmp = append(tmp, nums[len(nums)-k%l:]...)
	tmp = append(tmp, nums[:len(nums)-k%l]...)
	copy(nums, tmp)
}

func main() {

}
