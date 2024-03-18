package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestFunction(t *testing.T) {
	t.Parallel()
	type Input struct {
		Array  []int
		Target int
	}
	type TestCase struct {
		Question Input
		Answer   int
	}

	testCases := []TestCase{
		{
			Question: Input{
				Array:  []int{1, 3, 5, 6},
				Target: 5,
			},
			Answer: 2,
		},
		{
			Question: Input{
				Array:  []int{1, 3, 5, 6},
				Target: 2,
			},
			Answer: 1,
		},
		{
			Question: Input{
				Array:  []int{1, 3, 5, 6},
				Target: 7,
			},
			Answer: 4,
		},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.SearchInsert(tc.Question.Array, tc.Question.Target)
		if got != want {
			t.Errorf("want %v but got %v for test case %v, %d", want, got, tc.Question.Array, tc.Question.Target)
		}
	}
}
