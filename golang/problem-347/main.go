package main

import (
	"sort"
)

type kv [2]int

func Solve(nums []int, k int) []int {
	m := make(map[int]int)
	var ret []int

	for _, num := range nums {
		m[num]++
	}

	var kvs []kv
	for num, count := range m {
		kvs = append(kvs, kv{num, count})
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i][1] > kvs[j][1]
	})

	for i := 0; i < k; i++ {
		ret = append(ret, kvs[i][0])
	}

	return ret
}

func main() {

}
