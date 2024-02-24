package atoi_test

import (
	"atoi"
	"testing"
)

type TestCase struct {
	Test   string
	Result int
}

func TestAtoiConvertsStringsToNumbersCorrectly(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{"1", 1},
		{"29", 29},
		{"    -42", -42},
		{"42", 42},
		{"4193 with words", 4193},
		{"100000000000000000000000000000000", (1 << 31) - 1},
		{"-10000000000000000000000000000000", -(1 << 31)},
		{"just words", 0},
		{"-432 ok 34", -432},
		{"words and 987", 0},
		{"+-12", 0},
		{"2147483648", 2147483647},
		{"-2147483649", -2147483648},
	}

	for _, tc := range testCases {
		want := tc.Result
		got := atoi.MyAtoi(tc.Test)

		if want != got {
			t.Errorf("want %d, got %d", want, got)
		}
	}
}
