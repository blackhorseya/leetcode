package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(curr *ListNode) *ListNode {
	var prev *ListNode
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func Solve(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var (
		slow = head
		fast = head
	)

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow.Next = reverse(slow.Next)
	slow = slow.Next

	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}

	return true
}

func main() {

}
