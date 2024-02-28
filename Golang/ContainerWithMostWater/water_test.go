package waterarea_test

import (
	"testing"
	waterarea "water"
)

type TestCase struct {
	Test []int
	Want int
}

func TestMaxArea_ReturnsTheMaxAreaGivenArrayOfHeights(t *testing.T) {
	t.Parallel()

	testCases := []TestCase{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 5, 3, 6, 4, 5634, 5423, 123, 4, 43, 542, 243, 5432, 45}, 38024},
		{[]int{1, 1}, 1},
	}

	for no, tc := range testCases {
		want := tc.Want
		got := waterarea.MaxArea(tc.Test)

		if want != got {
			t.Errorf("want %d, got %d from test case no. %d", want, got, no+1)
		}
	}

}
