package leetcode_test

import (
	"encoding/json"
	"leetcode"
	"os"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question []int
		Answer   int
	}

	bigTestCase := []int{}
	file, err := os.Open("testdata/bigArray.json")
	if err != nil {
		t.Fatal(err)
	}
	err = json.NewDecoder(file).Decode(&bigTestCase)
	if err != nil {
		t.Fatal(err)
	}
	testCases := []TestCase{
		{Question: []int{0, 1}, Answer: 2},
		{Question: []int{0, 1, 0}, Answer: 2},
		{Question: bigTestCase, Answer: 44044},
	}

	for index, tc := range testCases {
		want := tc.Answer
		got := leetcode.FindMaxLength(tc.Question)

		if want != got {
			t.Errorf("want %d, got %d for test case no. %d", want, got, index+1)
		}
	}
}
