package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestFunction(t *testing.T) {
	t.Parallel()
	type Input struct {
		Array []int
		Dim   int
	}
	type TestCase struct {
		Question Input
		Answer   bool
	}

	testCases := []TestCase{
		{
			Question: Input{
				Array: []int{1, 2, 3, 1},
				Dim:   3,
			},
			Answer: true,
		},
		{
			Question: Input{
				Array: []int{1, 0, 1, 1},
				Dim:   1,
			},
			Answer: true,
		},
		{
			Question: Input{
				Array: []int{1, 2, 3, 1, 2, 3},
				Dim:   2,
			},
			Answer: false,
		},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.ContainsNearbyDuplicate(tc.Question.Array, tc.Question.Dim)
		if got != want {
			t.Errorf("want %v but got %v for test case %v, %d", want, got, tc.Question.Array, tc.Question.Dim)
		}
	}
}
