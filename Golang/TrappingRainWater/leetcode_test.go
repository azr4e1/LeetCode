package leetcode

import (
	"os"
	"testing"
)

func TestReturnCorrectAnswer(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Question []int
		Answer   int
	}

	testCases := []TestCase{
		{
			Question: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Answer:   6,
		},
		{
			Question: []int{4, 2, 0, 3, 2, 5},
			Answer:   9,
		},
	}

	for i, tc := range testCases {
		input := tc.Question
		Print(input, os.Stdout)
		want := tc.Answer
		got := trap(input)

		if want != got {
			t.Errorf("want %d, got %d for test case no. %d", want, got, i+1)
		}
	}
}
