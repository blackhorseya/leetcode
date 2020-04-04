package main

import (
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}

	var (
		tmp    = l
		result = ""
	)

	for tmp.Next != nil {
		result = result + strconv.Itoa(tmp.Val)
		tmp = tmp.Next
	}

	result = result + strconv.Itoa(tmp.Val)
	tmp = tmp.Next

	return result
}

func (l *ListNode) ToRevInt() int {
	var (
		tmp    = l
		result = 0
		index  = 0
	)

	for tmp.Next != nil {
		result += tmp.Val * int(math.Pow10(index))
		index++
		tmp = tmp.Next
	}

	result += tmp.Val * int(math.Pow10(index))
	index++
	tmp = tmp.Next

	return result
}

func IntToListNode(number int) *ListNode {
	var result = &ListNode{
		Val:  0,
		Next: nil,
	}

	if number == 0 {
		return result
	}

	var str = strconv.Itoa(number)
	result.Next = &ListNode{}
	var tmp = result.Next
	for i, s := range str {
		n, _ := strconv.Atoi(string(s))
		tmp.Val = n
		if i != len(str) - 1 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}

	result = result.Next
	return result
}

func Solve(first, second *ListNode) *ListNode {
	if first == nil || second == nil {
		return nil
	}

	sum := first.ToRevInt() + second.ToRevInt()
	result := IntToListNode(sum)
	return result
}

func main() {
	Solve(nil, nil)
}
