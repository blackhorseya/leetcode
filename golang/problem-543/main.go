package main

import (
	"leetcode/base"
)

func Solve(root *base.TreeNode) int {
	_, ret := maxDepth(root, 1)

	return ret - 1
}

func maxDepth(root *base.TreeNode, ret int) (int, int) {
	if root == nil {
		return 0, ret
	}

	left, lmax := maxDepth(root.Left, ret)
	right, rmax := maxDepth(root.Right, ret)

	ret = max(max(lmax, rmax), left+right+1)

	return max(left, right) + 1, ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func main() {
}
