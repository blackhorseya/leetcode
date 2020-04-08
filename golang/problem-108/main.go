package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	rootIndex := len(nums) / 2
	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  Solve(nums[:rootIndex]),
		Right: Solve(nums[rootIndex+1:]),
	}
}

func main() {

}
