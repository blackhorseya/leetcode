package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	var buff = &ListNode{}
	var result = buff
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			buff.Next = l2
			l2 = l2.Next
		} else {
			buff.Next = l1
			l1 = l1.Next
		}

		buff = buff.Next
	}

	if l1 == nil {
		buff.Next = l2
	}
	if l2 == nil {
		buff.Next = l1
	}

	return result.Next
}

func main() {

}
