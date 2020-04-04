package main

import (
	"sort"
)

func Solve(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
	return nums1
}

func main() {

}
