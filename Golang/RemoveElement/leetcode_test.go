package leetcode_test

import (
	"github.com/google/go-cmp/cmp"
	"leetcode"
	"testing"
)

type Input struct {
	Nums []int
	Val  int
}

type Output struct {
	Nums []int
	Val  int
}

func TestRemoveElement_RemovesElementsInPlaceAndReturnsLength(t *testing.T) {
	type TestCase struct {
		Question Input
		Answer   Output
	}
	testCases := []TestCase{
		{Question: Input{
			Nums: []int{3, 2, 2, 3},
			Val:  3,
		}, Answer: Output{
			Nums: []int{2, 2},
			Val:  2,
		}},
		{Question: Input{
			Nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:  2,
		}, Answer: Output{
			Nums: []int{0, 1, 4, 0, 3},
			Val:  5,
		},
		},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.RemoveElement(tc.Question.Nums, tc.Question.Val)
		if !cmp.Equal(want.Nums, tc.Question.Nums) {
			t.Error(cmp.Diff(want.Nums, tc.Question.Nums))
		}
		if got != want.Val {
			t.Errorf("want %d, got %d", want.Val, got)
		}
	}
}
