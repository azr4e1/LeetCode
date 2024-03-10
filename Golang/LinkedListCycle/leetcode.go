package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	pointers := map[*ListNode]bool{
		head: true,
	}
	for ; head.Next != nil; head = head.Next {
		_, ok := pointers[head.Next]
		if ok {
			return true
		}
		pointers[head.Next] = true
	}

	return false
}
