package leetcode_test

import (
	"github.com/google/go-cmp/cmp"
	"leetcode"
	"testing"
)

func TestLetterCombinations_ReturnsTheCorrectCombinationsGivenDigits(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question string
		Answer   []string
	}

	testCases := []TestCase{
		{Question: "23", Answer: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{Question: "", Answer: []string{}},
		{Question: "2", Answer: []string{"a", "b", "c"}},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.LetterCombinations(tc.Question)
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	}
}
