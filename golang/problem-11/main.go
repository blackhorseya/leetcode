package main

func Solve(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0

	for left < right {
		h := min(height[left], height[right])
		area = max(area, h*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {

}
