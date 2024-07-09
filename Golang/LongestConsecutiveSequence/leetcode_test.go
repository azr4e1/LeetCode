package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestReturnsCorrectResult(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Input  []int
		Answer int
	}

	testCases := []TestCase{
		{
			Input:  []int{100, 4, 200, 1, 3, 2},
			Answer: 4,
		},
		{
			Input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			Answer: 9,
		},
		{
			Input:  []int{},
			Answer: 0,
		},
	}

	for i, tc := range testCases {
		input := tc.Input
		want := tc.Answer
		got := leetcode.LongestConsecutive(input)
		if want != got {
			t.Errorf("want %d, got %d for test case no. %d", want, got, i+1)
		}
	}
}
