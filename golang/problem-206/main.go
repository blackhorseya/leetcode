package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var result = []int{head.Val}
	result = append(result, ToSlice(head.Next)...)

	return result
}

func Solve(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var result = &ListNode{
		Val:  0,
		Next: &ListNode{},
	}
	var tmp = result.Next
	slice := ToSlice(head)
	for index := len(slice) - 1; index >= 0; index-- {
		tmp.Val = slice[index]
		if index != 0 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}

	return result.Next
}

func main() {

}
