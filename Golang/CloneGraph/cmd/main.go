package main

import (
	"fmt"
	"leetcode"
)

func main() {
	node1 := leetcode.Node{Val: 1}
	node2 := leetcode.Node{Val: 2}
	node3 := leetcode.Node{Val: 3}
	node4 := leetcode.Node{Val: 4}
	node1.Neighbors = []*leetcode.Node{&node4, &node2}
	node2.Neighbors = []*leetcode.Node{&node1, &node3}
	node3.Neighbors = []*leetcode.Node{&node2, &node4}
	node4.Neighbors = []*leetcode.Node{&node1, &node3}

	fmt.Println(&node1)
	fmt.Println(leetcode.CloneGraph(&node1))
	fmt.Println(&node1)
}
