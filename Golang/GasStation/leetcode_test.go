package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestCanCompleteCircuit_ReturnsTheCorrectAnswer(t *testing.T) {
	t.Parallel()
	type Input struct {
		Gas  []int
		Cost []int
	}

	type TestCase struct {
		Question Input
		Answer   int
	}

	testCases := []TestCase{
		{
			Question: Input{
				Gas:  []int{1, 2, 3, 4, 5},
				Cost: []int{3, 4, 5, 1, 2},
			},
			Answer: 3,
		},
		{
			Question: Input{
				Gas:  []int{2, 3, 4},
				Cost: []int{3, 4, 3},
			},
			Answer: -1,
		},
		{
			Question: Input{
				Gas:  []int{5, 1, 2, 3, 4},
				Cost: []int{4, 4, 1, 5, 1},
			},
			Answer: 4,
		},
		{
			Question: Input{
				Gas:  []int{5, 8, 2, 8},
				Cost: []int{6, 5, 6, 6},
			},
			Answer: 3,
		},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.CanCompleteCircuit(tc.Question.Gas, tc.Question.Cost)
		if want != got {
			t.Errorf("want %d, got %d for test case %v", want, got, tc.Question)
		}
	}
}
