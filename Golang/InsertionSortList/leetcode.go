package leetcode

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	return &ListNode{
		Val:  input[0],
		Next: NewList(input[1:]),
	}
}

func (ln *ListNode) String() string {
	if ln == nil {
		return ""
	}
	var output = strconv.Itoa(ln.Val)
	for node := ln; node.Next != nil; node = node.Next {
		output += " " + strconv.Itoa(node.Next.Val)
	}

	return output
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for node := head; node.Next != nil; {
		if node.Next.Val < node.Val {
			misplaced := node.Next
			node.Next = misplaced.Next
			var prevNode *ListNode
			var node1 *ListNode
			for node1 = head; node1.Val < misplaced.Val; node1 = node1.Next {
				prevNode = node1
			}
			if prevNode == nil {
				head = misplaced
				head.Next = node1
			} else {
				prevNode.Next, misplaced.Next = misplaced, prevNode.Next
			}
			continue
		}
		node = node.Next
	}
	return head
}
