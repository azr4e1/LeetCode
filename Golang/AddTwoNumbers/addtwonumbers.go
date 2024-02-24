// You are given two non-empty linked lists representing two non-negative
// integers. The digits are stored in reverse order, and each of their nodes
// contains a single digit. Add the two numbers and return the sum as a linked
// list.
// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

type SumWithRest struct {
	Sum  int
	Rest int
}

func Add(x, y int) SumWithRest {
	sum := x + y
	rest := 0
	if sum >= 10 {
		sum -= 10
		rest = 1
	}

	return SumWithRest{sum, rest}
}

func NewList(input []int) *ListNode {
	if input == nil || len(input) == 0 {
		return &ListNode{}
	}
	if len(input) == 1 {
		return &ListNode{input[0], nil}
	}

	node := ListNode{
		Val:  input[0],
		Next: NewList(input[1:]),
	}
	return &node
}

func Decompose(x int) []int {
	if x == 0 {
		return []int{0}
	}

	s := []int{}
	for x > 0 {
		cur := x % 10
		s = append(s, cur)
		x = (x - cur) / 10
	}

	return s
}

func (ln *ListNode) Compose() int {
	var result int

	if ln.Next == nil {
		return ln.Val
	}

	result = 10*ln.Next.Compose() + ln.Val

	return result
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	newVal := Add(l1.Val, l2.Val)
	result := ListNode{Val: newVal.Sum}
	switch {
	case l1.Next == nil && l2.Next == nil:
		if newVal.Rest != 0 {
			result.Next = &ListNode{Val: newVal.Rest, Next: nil}
		}
		return &result
	case l1.Next == nil:
		l1.Next = &ListNode{}
	case l2.Next == nil:
		l2.Next = &ListNode{}
	}

	l1.Next.Val += newVal.Rest

	result.Next = AddTwoNumbers(l1.Next, l2.Next)

	return &result
}
