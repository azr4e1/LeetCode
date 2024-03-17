package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   int
	}

	testCases := []TestCase{
		{Question: []int{0, 1}, Answer: 2},
		{Question: []int{0, 1, 0}, Answer: 2},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.FindMaxLength(tc.Question)

		if want != got {
			t.Errorf("want %d, got %d for test case %v", want, got, tc.Question)
		}
	}
}
