package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestThreeSumClosest_CorrectlyFindsTheThreeNumbersClosestToTarget(t *testing.T) {
	t.Parallel()
	inputArray := []int{-1, 2, 1, -4}
	inputTarget := 1
	want := 2
	got := leetcode.ThreeSumClosest(inputArray, inputTarget)
	if got != want {
		t.Errorf("want %d, got %d for test case %v, %d", want, got, inputArray, inputTarget)
	}
	inputArray = []int{0, 0, 0}
	inputTarget = 1
	want = 0
	got = leetcode.ThreeSumClosest(inputArray, inputTarget)
	if got != want {
		t.Errorf("want %d, got %d for test case %v, %d", want, got, inputArray, inputTarget)
	}
	inputArray = []int{3, 4, 2, 1, 2, 43, 2, 1, 43, 31, 34, 2, 43}
	inputTarget = 12
	want = 12
	got = leetcode.ThreeSumClosest(inputArray, inputTarget)
	if got != want {
		t.Errorf("want %d, got %d for test case %v, %d", want, got, inputArray, inputTarget)
	}
}
