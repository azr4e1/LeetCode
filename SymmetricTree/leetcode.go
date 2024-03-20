package leetcode

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirroredTree(root.Left, root.Right)
}

func isMirroredTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	isSameLeft := isMirroredTree(p.Left, q.Right)
	isSameRight := isMirroredTree(p.Right, q.Left)

	return isSameLeft && isSameRight
}
