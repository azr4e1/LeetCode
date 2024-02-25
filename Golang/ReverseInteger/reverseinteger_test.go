package reverseinteger_test

import (
	ri "reverseinteger"
	"testing"
)

type TestCase struct {
	Test int
	Want int
}

func TestReverseReturnsTheCorrectInvertedNumber(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{5434563, 3654345},
	}

	for _, tc := range testCases {
		want := tc.Want
		got := ri.Reverse(tc.Test)
		if got != want {
			t.Errorf("want %d, got %d for test case %d", want, got, tc.Test)
		}
	}
}

func TestReverseReturnsZeroIfOverflow(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{2147483647, 0},
		{-2147483648, 0},
		{2147483643, 0},
		{-1000000003, 0},
	}

	for _, tc := range testCases {
		want := tc.Want
		got := ri.Reverse(tc.Test)
		if got != want {
			t.Errorf("want %d, got %d for test case %d where reverse overflows", want, got, tc.Test)
		}
	}
}
