package leetcode_test

import (
	"github.com/google/go-cmp/cmp"
	"leetcode"
	"testing"
)

func TestThreeSum_CorrectlyGivesAllTheTriplets(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   [][]int
	}

	testCases := []TestCase{
		{Question: []int{-1, 0, 1, 2, -1, -4}, Answer: [][]int{
			{-1, 0, 1},
			{-1, -1, 2},
		},
		},
		{Question: []int{0, 1, 1}, Answer: [][]int{}},
		{Question: []int{0, 0, 0}, Answer: [][]int{{0, 0, 0}}},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.ThreeSum(tc.Question)

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	}
}

func TestTwoSum_ReturnsCorrectPairs(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Input1 []int
		Input2 int
		Answer [][]int
	}
	testCases := []TestCase{
		{Input1: []int{1, 2, 3, 4}, Input2: 6, Answer: [][]int{{2, 4}}},
		{Input1: []int{1, 2, 3, 4}, Input2: 5, Answer: [][]int{{2, 3}, {1, 4}}},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.TwoSum(tc.Input1, tc.Input2)

		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	}
}
