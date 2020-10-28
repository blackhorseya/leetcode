package main

import "leetcode/base"

func Solve(t1 *base.TreeNode, t2 *base.TreeNode) *base.TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	ret := &base.TreeNode{Val: t1.Val + t2.Val}
	ret.Left = Solve(t1.Left, t2.Left)
	ret.Right = Solve(t1.Right, t2.Right)

	return ret
}

func main() {

}
