package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	node := ListNode{
		Val:  input[0],
		Next: NewList(input[1:]),
	}

	return &node
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, small, large *ListNode
	if list1.Val < list2.Val {
		small = list1
		large = list2
	} else {
		small = list2
		large = list1
	}
	for head = small; small.Next != nil && small.Next.Val < large.Val; small = small.Next {
	}
	small.Next = MergeTwoLists(small.Next, large)

	return head
}
