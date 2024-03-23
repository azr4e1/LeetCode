package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	var el *TreeNode
	cache := make(map[int]struct{})
	values := []int{}
	var minimum = 100000

	for len(queue) > 0 {
		el, queue = queue[0], queue[1:]
		if el == nil {
			continue
		}
		if _, ok := cache[el.Val]; !ok {
			cache[el.Val] = struct{}{}
			for _, val := range values {
				diff := Abs(val - el.Val)
				if diff < minimum {
					minimum = diff
				}
			}
			values = append(values, el.Val)
		}
		queue = append(queue, el.Left, el.Right)
	}

	return minimum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
