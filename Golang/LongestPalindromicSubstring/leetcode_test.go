package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestLongestPalindrome_ReturnsTheLongestPalindromicSubstring(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question string
		Answer   string
	}

	testCases := []TestCase{
		{Question: "babad", Answer: "bab"},
		{Question: "cbbd", Answer: "bb"},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.LongestPalindrome(tc.Question)
		if want != got {
			t.Errorf("want %q, got %q from test case %q", want, got, tc.Question)
		}
	}
}
