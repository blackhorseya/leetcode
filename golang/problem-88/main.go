package main

func Solve(nums1 []int, m int, nums2 []int, n int) {
	var (
		merged = make([]int, len(nums1))
		i      = 0
		j      = 0
		k      = 0
	)
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}

		k++
	}

	for i < m {
		merged[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		merged[k] = nums2[j]
		j++
		k++
	}

	copy(nums1, merged)
}

func main() {

}
