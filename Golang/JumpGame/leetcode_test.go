package leetcode_test

import (
	"encoding/json"
	"leetcode"
	"os"
	"testing"
)

func TestCanJump_ReturnsIfWeCanReachTheLastIndex(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   bool
	}

	testCases := []TestCase{
		{Question: []int{2, 3, 1, 1, 4}, Answer: true},
		{Question: []int{3, 2, 1, 0, 4}, Answer: false},
	}
	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.CanJump(tc.Question)

		if got != want {
			t.Errorf("want %v, got %v for test case %v", want, got, tc.Question)
		}
	}
	input := []int{}
	file, err := os.Open("./testdata/bigInput.json")
	if err != nil {
		t.Fatal(err)
	}
	json.NewDecoder(file).Decode(&input)
	want := false
	got := leetcode.CanJump(input)

	if got != want {
		t.Errorf("want %v, got %v for very big input", want, got)
	}
}
