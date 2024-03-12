package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestMajorityElement_ReturnsTheCorrectResult(t *testing.T) {
	type TestCase struct {
		Question []int
		Answer   int
	}

	testCases := []TestCase{
		{Question: []int{3, 2, 3}, Answer: 3},
		{Question: []int{2, 2, 1, 1, 1, 2, 2}, Answer: 2},
		{Question: []int{-1, 1, 1, 1, 2, 1}, Answer: 1},
		{Question: []int{1, 1, 1, 2, 2, 2, 3, 1}, Answer: 1},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.MajorityElement(tc.Question)
		if want != got {
			t.Errorf("want %d, got %d from test case %v", want, got, tc.Question)
		}
	}
}
