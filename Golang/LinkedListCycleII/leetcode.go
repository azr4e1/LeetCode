package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pointers := map[*ListNode]bool{
		head: true,
	}
	for ; head.Next != nil; head = head.Next {
		_, ok := pointers[head.Next]
		if ok {
			return head.Next
		}
		pointers[head.Next] = true
	}

	return nil

}
