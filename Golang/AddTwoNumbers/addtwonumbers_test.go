package addtwonumbers_test

import (
	atn "addtwonumbers"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewList_CreatesASinglyLinkedList(t *testing.T) {
	t.Parallel()
	input := []int{3, 2, 4, 12, 32, 12, 4, 76, 587, 10}
	listNodes := make([]atn.ListNode, 10, 10)
	for i := 0; i < len(input); i++ {
		listNodes[i] = atn.ListNode{
			Val:  input[i],
			Next: nil,
		}
	}
	for i := 0; i < len(listNodes)-1; i++ {
		listNodes[i].Next = &listNodes[i+1]
	}
	want := &(listNodes[0])

	got := atn.NewList(input)

	if !cmp.Equal(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}
func TestNewList_ReturnsEmptyNodeIfInputIsInvalid(t *testing.T) {
	t.Parallel()
	got := atn.NewList(nil)
	if !cmp.Equal(*got, atn.ListNode{}) {
		t.Error("want empty node for nil input, got nil")
	}

	got = atn.NewList([]int{})
	if !cmp.Equal(*got, atn.ListNode{}) {
		t.Error("want empty node for empty slice, got nil")
	}
}

func TestComposeReturnsIntegerRepresentedByList(t *testing.T) {
	t.Parallel()
	want := 524535
	list := atn.NewList([]int{5, 3, 5, 4, 2, 5})
	got := list.Compose()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

	list = &atn.ListNode{}
	got = list.Compose()

	if 0 != got {
		t.Errorf("want 0, got %d", got)
	}
}

func TestDecomposeTurnsIntegerIntoInversedSequenceOfDigits(t *testing.T) {
	t.Parallel()
	input := 43252341
	want := []int{1, 4, 3, 2, 5, 2, 3, 4}
	got := atn.Decompose(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	input = 0
	want = []int{0}
	got = atn.Decompose(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestAddTwoNumbersAddsTwoInvertedSinglyLinkedListCorrectly(t *testing.T) {
	t.Parallel()
	list1 := atn.NewList([]int{2, 3, 1, 4, 2, 4})
	list2 := atn.NewList([]int{3, 2, 1, 5, 4, 6})

	want := atn.NewList([]int{5, 5, 2, 9, 6, 0, 1})
	got := atn.AddTwoNumbers(list1, list2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	list1 = atn.NewList([]int{2, 3, 1, 4, 2, 4, 8, 9})
	list2 = atn.NewList([]int{3, 2, 1, 5, 4, 6})

	want = atn.NewList([]int{5, 5, 2, 9, 6, 0, 9, 9})
	got = atn.AddTwoNumbers(list1, list2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	list1 = atn.NewList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	list2 = atn.NewList([]int{5, 6, 4})

	want = atn.NewList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	got = atn.AddTwoNumbers(list1, list2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddCreatesASumStructWithCorrectRest(t *testing.T) {
	t.Parallel()
	x, y := 6, 9
	want := atn.SumWithRest{5, 1}
	got := atn.Add(x, y)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	x, y = 3, 2
	want = atn.SumWithRest{5, 0}
	got = atn.Add(x, y)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
