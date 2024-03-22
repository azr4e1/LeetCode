package leetcode

import (
	// "fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeLevel struct {
	Node  *TreeNode
	Level int
}

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}

	queue := []*TreeLevel{
		{root.Left, 1},
		{root.Right, 1},
	}
	storage := make(map[int][]int)
	storage[0] = []int{root.Val}

	var el *TreeLevel
	length := 0

	for len(queue) > 0 {
		el, queue = queue[0], queue[1:]
		if el.Node == nil {
			continue
		}
		node, level := el.Node, el.Level
		if level > length {
			length = level
		}
		storage[level] = append(storage[level], node.Val)
		queue = append(queue, &TreeLevel{node.Left, level + 1}, &TreeLevel{node.Right, level + 1})
	}
	avgLevels := make([]float64, length+1)
	for index, vals := range storage {
		avgLevels[index] = Mean(vals)
	}

	return avgLevels
}

func Mean(vals []int) float64 {
	if vals == nil || len(vals) == 0 {
		return math.NaN()
	}
	length := float64(len(vals))
	sum := 0.0
	for _, val := range vals {
		sum += float64(val)
	}

	return sum / length
}
