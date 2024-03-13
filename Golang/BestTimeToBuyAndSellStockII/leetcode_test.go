package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestMaxProfit_ReturnsTheCorrectProfit(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   int
	}
	testCases := []TestCase{
		{Question: []int{7, 1, 5, 3, 6, 4}, Answer: 7},
		{Question: []int{1, 2, 3, 4, 5}, Answer: 4},
		{Question: []int{7, 6, 4, 3, 1}, Answer: 0},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.MaxProfit(tc.Question)

		if want != got {
			t.Errorf("want %d, got %d for test case %v", want, got, tc.Question)
		}
	}
}
