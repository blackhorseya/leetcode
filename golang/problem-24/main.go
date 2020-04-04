package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if array := ToSlice(l); array == nil {
		return ""
	} else {
		return strings.Join(array, "->")
	}
}

func ToSlice(l *ListNode) []string {
	if l == nil {
		return nil
	}

	var result []string
	var tmp = l
	for tmp.Next != nil {
		result = append(result, strconv.Itoa(tmp.Val))
		tmp = tmp.Next
	}

	return append(result, strconv.Itoa(tmp.Val))
}

func Solve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head.Next.Next = Solve(head.Next.Next)
	c := head
	head = head.Next
	c.Next = head.Next
	head.Next = c
	return head
}

func main() {
	
}
