package leetcode

import (
	"errors"
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

type Queue []*Node

func (q *Queue) Push(el *Node) {
	*q = append(*q, el)
}

func (q *Queue) Pop() (*Node, error) {
	var el *Node
	if len(*q) == 0 {
		return el, errors.New("Queue is empty")
	}
	*q, el = (*q)[1:], (*q)[0]

	return el, nil
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (n *Node) String() string {
	q := Queue{n}
	explored := map[*Node]bool{n: true}
	output := "["
	for !q.IsEmpty() {
		v, err := q.Pop()
		if err != nil {
			panic(err)
		}
		output += fmt.Sprintf("%p, %d: [", v, v.Val)
		for _, newNode := range v.Neighbors {
			output += fmt.Sprintf("%d, ", newNode.Val)
			_, ok := explored[newNode]
			if ok {
				continue
			}
			explored[newNode] = true
			q.Push(newNode)
		}
		output += "] "
	}
	output += "]"

	return output
}

func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	var newInitialNode = &Node{Val: node.Val}
	var newElements = Queue{node}
	var copiedElements = Queue{newInitialNode}
	var exploredElements = map[*Node]*Node{node: newInitialNode}
	for !newElements.IsEmpty() {
		v, err := newElements.Pop()
		if err != nil {
			panic(err)
		}
		copiedV, copiedErr := copiedElements.Pop()
		if copiedErr != nil {
			panic(copiedErr)
		}
		for _, newNode := range v.Neighbors {
			oldCopiedNode, ok := exploredElements[newNode]
			if ok {
				copiedV.Neighbors = append(copiedV.Neighbors, oldCopiedNode)
				continue
			}
			newCopiedNode := &Node{Val: newNode.Val}
			copiedV.Neighbors = append(copiedV.Neighbors, newCopiedNode)

			exploredElements[newNode] = newCopiedNode
			newElements.Push(newNode)
			copiedElements.Push(newCopiedNode)
		}
	}

	return newInitialNode
}
