package main

import (
	"sort"
)

func Solve(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	ret := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		last := ret[len(ret)-1]

		if include(interval, last) {
			ret[len(ret)-1] = []int{last[0], max(last[1], interval[1])}
		} else {
			ret = append(ret, interval)
		}
	}

	return ret
}

func include(a, b []int) bool {
	cur1, cur2 := a[0], a[1]
	next1, next2 := b[0], b[1]

	if cur1 <= next2 && next1 <= cur2 {
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

}
