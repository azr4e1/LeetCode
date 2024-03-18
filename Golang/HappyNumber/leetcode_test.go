package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestIsHappy_ReturnsTrueIfNumberIsHappy(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question int
		Answer   bool
	}

	testCases := []TestCase{
		{Question: 19, Answer: true},
		{Question: 2, Answer: false},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.IsHappy(tc.Question)
		if want != got {
			t.Errorf("want %v, got %v for test case %d", want, got, tc.Question)
		}
	}
}
