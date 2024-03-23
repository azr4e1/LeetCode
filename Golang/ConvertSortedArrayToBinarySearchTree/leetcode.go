// Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	middlePoint := len(nums) / 2

	return &TreeNode{
		Val:   nums[middlePoint],
		Left:  sortedArrayToBST(nums[:middlePoint]),
		Right: sortedArrayToBST(nums[middlePoint+1:]),
	}
}
