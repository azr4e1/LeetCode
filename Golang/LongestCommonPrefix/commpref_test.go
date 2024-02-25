package commpref_test

import (
	"commpref"
	"testing"
)

type TestCase struct {
	Test []string
	Want string
}

func TestLongestCommonPrefixReturnsCommonPrefixes(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"aac", "acab", "aa", "abba", "aa"}, "a"},
	}
	for _, tc := range testCases {
		want := tc.Want
		got := commpref.LongestCommonPrefix(tc.Test)
		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	}
}
