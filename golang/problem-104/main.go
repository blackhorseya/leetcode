package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftDep, rightDep = Solve(root.Left), Solve(root.Right)
	if leftDep >= rightDep {
		return leftDep + 1
	} else {
		return rightDep + 1
	}
}

func main() {

}
