package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestHIndex_CalculatesTheCorrectHIndexOfTheGivenResearchers(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   int
	}

	testCases := []TestCase{
		{Question: []int{3, 0, 6, 1, 5}, Answer: 3},
		{Question: []int{1, 3, 1}, Answer: 1},
		{Question: []int{3, 3}, Answer: 2},
		{Question: []int{4, 4, 4}, Answer: 3},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.HIndex(tc.Question)

		if got != want {
			t.Errorf("want %v, got %v for test case %v", want, got, tc.Question)
		}
	}
}
