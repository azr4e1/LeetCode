package palindrome_test

import (
	"palindrome"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type TestCase struct {
	Test     int
	Response bool
}

func TestDecomposeTurnsIntegerIntoSliceOfSingleDigitsIntegers(t *testing.T) {
	t.Parallel()
	input := 43252352
	want := []int{2, 5, 3, 2, 5, 2, 3, 4}
	got := palindrome.Decompose(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	input = 0
	want = []int{0}
	got = palindrome.Decompose(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	input = -534563
	want = []int{}
	got = palindrome.Decompose(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
func TestComposeTurnsSliceOfSingleDigitIntegersIntoNumber(t *testing.T) {
	t.Parallel()
	input := []int{2, 5, 3, 2, 5, 2, 3, 4}
	want := 43252352
	got, err := palindrome.Compose(input)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestComposeReturnsErrorForInvalidInput(t *testing.T) {
	input := []int{4, 32, 2, 5, 2, 3, 5, 2}
	_, err := palindrome.Compose(input)
	if err == nil {
		t.Error("expected error, got nil for slice of non-single digit integers")
	}

	input = []int{4, 32, 2, 5, -2, 3, -5, 2}
	_, err = palindrome.Compose(input)
	if err == nil {
		t.Error("expected error, got nil for slice of non-positive integers")
	}
}

func TestIsPalindrome_ReturnsTheCorrectAnswer(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{121, true},
		{-121, false},
		{10, false},
		{3245235, false},
		{10101, true},
		{1111, true},
		{133141331, true},
		{1441, true},
		{98765346432, false},
	}
	for _, tc := range testCases {
		want := tc.Response
		got := palindrome.IsPalindrome(tc.Test)

		if want != got {
			t.Errorf("want %t, got %t for test case %d", want, got, tc.Test)
		}
	}
}
