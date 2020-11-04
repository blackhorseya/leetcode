package main

func Solve(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right, max := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {

}
