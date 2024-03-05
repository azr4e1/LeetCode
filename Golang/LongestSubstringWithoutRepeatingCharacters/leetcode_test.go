package leetcode_test

import (
	"leetcode"
	"testing"
)

type TestCase struct {
	Question string
	Answer   int
}

func TestLengthOfLongestSubstring_FindsTheLengthOfLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{Question: "abcabcbb", Answer: 3},
		{Question: "bbbbb", Answer: 1},
		{Question: "pwwkew", Answer: 3},
		{Question: "aab", Answer: 2},
		{Question: "aabaab!bb", Answer: 3},
		{Question: "museuwzbczdejn", Answer: 7},
		{Question: "ohomm", Answer: 3},
		{Question: "eeydgwdykpv", Answer: 7},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.LengthOfLongestSubstring(tc.Question)
		if want != got {
			t.Errorf("want %d, got %d from test case %q", want, got, tc.Question)
		}
	}
}
