package leetcode_test

import (
	"github.com/google/go-cmp/cmp"
	"leetcode"
	"testing"
)

type Input struct {
	Num1 []int
	M    int
	Num2 []int
	N    int
}

func TestMergeMergedCorrectlyTheTwoArrays(t *testing.T) {
	type TestCase struct {
		Question *Input
		Answer   []int
	}
	t.Parallel()

	testCases := []TestCase{
		{Question: &Input{
			Num1: []int{1, 2, 3, 0, 0, 0},
			M:    3,
			Num2: []int{2, 5, 6},
			N:    3,
		}, Answer: []int{1, 2, 2, 3, 5, 6}},
		{Question: &Input{
			Num1: []int{1},
			M:    1,
			Num2: []int{},
			N:    0,
		}, Answer: []int{1}},
		{Question: &Input{
			Num1: []int{0},
			M:    0,
			Num2: []int{1},
			N:    1,
		}, Answer: []int{1}},
	}

	for _, tc := range testCases {
		want := tc.Answer
		leetcode.Merge(tc.Question.Num1, tc.Question.Num2, tc.Question.M, tc.Question.N)

		if !cmp.Equal(tc.Question.Num1, want) {
			t.Error(cmp.Diff(tc.Question.Num1, want))
		}
	}

}
