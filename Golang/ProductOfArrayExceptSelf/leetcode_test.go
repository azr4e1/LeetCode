package leetcode_test

import (
	"leetcode"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProductExceptSelf_ReturnsArrayOfProducts(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Input  []int
		Output []int
	}
	testCases := []TestCase{
		{Input: []int{1, 2, 3, 4}, Output: []int{24, 12, 8, 6}},
		{Input: []int{-1, 1, 0, -3, 3}, Output: []int{0, 0, 9, 0, 0}},
		{Input: []int{0, 2, 3, 1, 4, 0, 5}, Output: []int{0, 0, 0, 0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		want := tc.Output
		got := leetcode.ProductExceptSelf(tc.Input)
		if !cmp.Equal(want, got) {
			t.Error(cmp.Diff(want, got))
		}
	}
}
