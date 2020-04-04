package main

type ListNode struct {
    Val int
    Next *ListNode
}

func Solve(head *ListNode) *ListNode {
	var cur = head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}

func main() {
	
}
