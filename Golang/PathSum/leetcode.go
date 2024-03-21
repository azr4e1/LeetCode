// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

// A leaf is a node with no children.
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	remainingTarget := targetSum - root.Val
	if root.Left == nil && root.Right == nil && remainingTarget == 0 {
		return true
	}
	var left, right bool
	if root.Left != nil {
		left = hasPathSum(root.Left, remainingTarget)
	}
	if root.Right != nil {
		right = hasPathSum(root.Right, remainingTarget)
	}
	
	return left || right
}

