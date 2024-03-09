package leetcode_test

import (
	"leetcode"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValid_ReturnsValidity(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		Question string
		Answer   bool
	}

	testCases := []TestCase{
		{Question: "()", Answer: true},
		{Question: "()[]{}", Answer: true},
		{Question: "(]", Answer: false},
		{Question: "([])", Answer: true},
		{Question: "([{)}]", Answer: false},
		{Question: "[", Answer: false},
		{Question: "}", Answer: false},
	}

	for _, tc := range testCases {
		want := tc.Answer
		got := leetcode.IsValid(tc.Question)

		if got != want {
			t.Errorf("want %v, got %v from testcase %q", want, got, tc.Question)
		}
	}
}

func TestQueueAdd_AddsAnElementAtTheEndOfTheQueue(t *testing.T) {
	t.Parallel()
	input := []byte("(){}")
	queue := leetcode.Queue(input)
	queue.Add(']')
	newVal := append(input, ']')

	if !cmp.Equal([]byte(queue), newVal) {
		t.Error(cmp.Diff([]byte(queue), newVal))
	}
}

func TestQueuePop_RemovesAnElementFromEndOfQueue(t *testing.T) {
	t.Parallel()
	input := []byte("(){}")
	queue := leetcode.Queue(input)
	el, err := queue.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if el != '}' {
		t.Errorf("want %s, got %s", "}", string(el))
	}
	if !cmp.Equal([]byte(queue), []byte("(){")) {
		t.Error(cmp.Diff([]byte(queue), []byte("(){")))
	}
}

func TestQueuePop_WhenQueueIsEmptyReturnsError(t *testing.T) {
	t.Parallel()
	input := []byte{}
	queue := leetcode.Queue(input)
	_, err := queue.Pop()
	if err == nil {
		t.Fatal("want err, got nil")
	}
}
