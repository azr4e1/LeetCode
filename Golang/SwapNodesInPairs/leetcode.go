package leetcode

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	s := ""
	for l.Next != nil {
		s = s + " " + strconv.Itoa(l.Val)
		l = *(l.Next)
	}
	s = s + " " + strconv.Itoa(l.Val)

	return s
}

func FromSlice(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	if len(s) == 1 {
		return &ListNode{s[0], nil}
	}
	node := ListNode{s[0], FromSlice(s[1:])}
	return &node
}

func FromList(l *ListNode) []int {
	result := []int{l.Val}
	if l.Next == nil {
		return result
	}
	partialRes := FromList(l.Next)
	for _, el := range partialRes {
		result = append(result, el)
	}

	return result
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		newHead := head.Next
		head.Next, newHead.Next = nil, head
		return newHead
	}

	if head.Next.Next.Next == nil {
		newHead := head.Next
		head.Next, newHead.Next = newHead.Next, head
		return newHead
	}

	newHead := head.Next
	thirdNode := newHead.Next
	lastNode := thirdNode.Next
	head.Next, newHead.Next = lastNode, head
	lastNode.Next, thirdNode.Next = thirdNode, SwapPairs(lastNode.Next)

	return newHead
}
