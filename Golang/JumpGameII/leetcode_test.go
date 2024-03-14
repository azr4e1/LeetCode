package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestCanJump_ReturnsIfWeCanReachTheLastIndex(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   int
	}

	testCases := []TestCase{
		{Question: []int{2, 3, 1, 1, 4}, Answer: 2},
		{Question: []int{2, 3, 0, 1, 4}, Answer: 2},
		{Question: []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}, Answer: 5},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.Jump(tc.Question)

		if got != want {
			t.Errorf("want %v, got %v for test case %v", want, got, tc.Question)
		}
	}
}
