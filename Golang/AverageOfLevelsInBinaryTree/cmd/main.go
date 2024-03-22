package main

import (
	"fmt"
	lc "leetcode"
)

func main() {
	node := lc.TreeNode{
		Val: 3,
		Left: &lc.TreeNode{
			Val: 9,
		},
		Right: &lc.TreeNode{
			Val: 20,
			Left: &lc.TreeNode{
				Val: 15,
			},
			Right: &lc.TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(lc.AverageOfLevels(&node))

	node2 := lc.TreeNode{
		Val: 5,
		Left: &lc.TreeNode{
			Val: 2,
		},
		Right: &lc.TreeNode{
			Val: -3,
		},
	}
	fmt.Println(lc.AverageOfLevels(&node2))
}
