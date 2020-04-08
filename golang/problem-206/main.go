package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(head *ListNode) *ListNode {
	var result *ListNode

	for head != nil {
		current := head
		head = head.Next

		current.Next = result
		result = current
	}

	return result
}

func main() {

}
