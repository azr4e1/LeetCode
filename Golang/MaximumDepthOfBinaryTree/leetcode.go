// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
package leetcode

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + maxDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + maxDepth(root.Left)
	}

	depthR, depthL := 1+maxDepth(root.Right), 1+maxDepth(root.Left)
	if depthR > depthL {
		return depthR
	}
	return depthL
}
