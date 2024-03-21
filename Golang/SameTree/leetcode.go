package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	isSameLeft := isSameTree(p.Left, q.Left)
	isSameRight := isSameTree(p.Right, q.Right)

	return isSameLeft && isSameRight
}
