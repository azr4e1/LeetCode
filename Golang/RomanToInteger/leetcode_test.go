package leetcode_test

import (
	"leetcode"
	"testing"
)

func TestRomanToInt_ConvertsRomanNumeralToInt(t *testing.T) {
	t.Parallel()
	var want int
	var input string
	var got int

	input = "III"
	want = 3
	got = leetcode.RomanToInt(input)
	if want != got {
		t.Errorf("want %d, got %d for testcase %s", want, got, input)
	}

	input = "LVIII"
	want = 58
	got = leetcode.RomanToInt(input)
	if want != got {
		t.Errorf("want %d, got %d for testcase %s", want, got, input)
	}

	input = "MCMXCIV"
	want = 1994
	got = leetcode.RomanToInt(input)
	if want != got {
		t.Errorf("want %d, got %d for testcase %s", want, got, input)
	}
}
