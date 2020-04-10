package main

type ListNode struct {
    Val int
    Next *ListNode
}

func Solve(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for {
		if fast == nil || fast.Next == nil {
			return false
		}

		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}

func main() {
	
}
