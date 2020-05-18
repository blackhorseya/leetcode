package main

import "sort"

func Solve(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return Solve(nums2, nums1)
	}

	var ret []int
	for _, num := range nums2 {
		contain, pos := contain(nums1, num)
		if contain {
			ret = append(ret, num)
			nums1 = append(nums1[:pos], nums1[pos+1:]...)
		}
	}

	sort.Ints(ret)
	return ret
}

func contain(nums []int, target int) (contain bool, pos int) {
	for index, num := range nums {
		if num == target {
			return true, index
		}
	}

	return false, 0
}

func main() {

}
