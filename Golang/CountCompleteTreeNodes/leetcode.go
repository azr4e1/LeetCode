package leetcode

import "errors"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Vals []*TreeNode
}

func (q *Queue) Add(el ...*TreeNode) {
	q.Vals = append(q.Vals, el...)
}

func (q *Queue) Pop() (*TreeNode, error) {
	if len(q.Vals) == 0 {
		return nil, errors.New("Queue is empty")
	}
	el, newQueue := q.Vals[0], q.Vals[1:]
	q.Vals = newQueue
	return el, nil
}

func (q *Queue) Empty() bool {
	return len(q.Vals) == 0
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	counter := 1
	queue := Queue{Vals: []*TreeNode{root.Left, root.Right}}
	for !queue.Empty() {
		el, err := queue.Pop()
		if err != nil {
			panic(err)
		}
		if el == nil {
			continue
		}
		counter++
		queue.Add(el.Left, el.Right)
	}
	return counter
}
