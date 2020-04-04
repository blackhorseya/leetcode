package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if array := RetrievedListNode(l); array == nil {
		return ""
	} else {
		var result = ""
		for _, num := range array {
			result += strconv.Itoa(num)
		}

		return result
	}
}

func RetrievedListNode(l *ListNode) []int {
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
	if first == nil || second == nil {
		return nil
	}

	slice1 := RetrievedListNode(first)
	slice2 := RetrievedListNode(second)

	if len(slice2) > len(slice1) {
		var tmp = slice1
		slice1 = slice2
		slice2 = tmp
	}

	var sumSlice []int
	var over = 0
	for index := 0; index < len(slice1); index++ {
		var sum = 0
		if over > 0 {
			sum += over
			over = 0
		}

		if index >= len(slice2) {
			sum += slice1[index]
		} else {
			sum += slice1[index] + slice2[index]
		}
		if sum / 10 >= 1 {
			sum = sum % 10
			over++
		}

		sumSlice = append(sumSlice, sum)
	}
	if over > 0 {
		sumSlice = append(sumSlice, over)
	}

	var result = &ListNode{
		Val:  0,
		Next: &ListNode{},
	}
	var tmp = result.Next
	for index, num := range sumSlice {
		tmp.Val = num
		if index != len(sumSlice) - 1 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}

	return result.Next
}

func main() {
	Solve(nil, nil)
}
