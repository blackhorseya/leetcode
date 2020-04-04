package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		} else {
			return Solve(p.Left, q.Left) && Solve(p.Right, q.Right)
		}
	}

	if p == nil && q == nil {
		return true
	}

	return false
}

func main() {

}
