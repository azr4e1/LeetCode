package leetcode_test

import (
	"github.com/google/go-cmp/cmp"
	"leetcode"
	"testing"
)

func TestGenerateParenthesis_GeneratesAllCombinationsOfValidPairs(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question int
		Answer   []string
	}

	testCases := []TestCase{
		{Question: 3, Answer: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{Question: 2, Answer: []string{"(())", "()()"}},
		{Question: 1, Answer: []string{"()"}},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.GenerateParenthesis(tc.Question)
		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(got, want))
		}
	}
}
