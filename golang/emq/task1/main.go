package main

func Solution(blocks []int) int {
	length := len(blocks)

	l := make([]int, length)
	for start := 1; start < length; start++ {
		if blocks[start] <= blocks[start-1] {
			l[start] = l[start-1] + 1
		} else {
			l[start] = 0
		}
	}

	r := make([]int, length)
	for start := length - 2; start >= 0; start-- {
		if blocks[start] <= blocks[start+1] {
			r[start] = r[start+1] + 1
		} else {
			r[start] = 0
		}
	}

	ret := 0
	for i := range l {
		ret = max(ret, l[i] + r[i] + 1)
	}

	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {

}
