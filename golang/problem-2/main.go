package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if array := ToSlice(l); array == nil {
		return ""
	} else {
		var result = ""
		for _, num := range array {
			result += strconv.Itoa(num)
		}

		return result
	}
}

func ToSlice(l *ListNode) []int {
	if l == nil {
		return nil
	}

	var result []int
	var tmp = l
	for tmp.Next != nil {
		result = append(result, tmp.Val)
		tmp = tmp.Next
	}

	return append(result, tmp.Val)
}

func Solve(first, second *ListNode) *ListNode {
	var (
		result = &ListNode{
			Val:  0,
			Next: &ListNode{},
		}
		carry = 0
		tmp   = result
	)
	for first != nil || second != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if first != nil {
			v1 = first.Val
			first = first.Next
		}
		if second != nil {
			v2 = second.Val
			second = second.Next
		}

		val := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		tmp.Next = &ListNode{
			Val: val,
		}
		tmp = tmp.Next
	}

	return result.Next
}

func main() {
	Solve(nil, nil)
}
