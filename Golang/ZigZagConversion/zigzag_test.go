package zigzag_test

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"zigzag"
)

type Input struct {
	String string
	Rows   int
}

func TestToMatrix_ConvertsStringIntoItsZigZagMatrixRepresentation(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Test   Input
		Result [][]byte
	}
	testCases := []TestCase{
		{
			Test: Input{"PAYPALISHIRING", 4},
			Result: [][]byte{
				{'P', 'I', 'N'},
				{'A', 'L', 'S', 'I', 'G'},
				{'Y', 'A', 'H', 'R'},
				{'P', 'I'},
			},
		},
		{
			Test: Input{"PAYPALISHIRING", 3},
			Result: [][]byte{
				{'P', 'A', 'H', 'N'},
				{'A', 'P', 'L', 'S', 'I', 'I', 'G'},
				{'Y', 'I', 'R'},
			},
		},
		{
			Test: Input{"A", 1},
			Result: [][]byte{
				{'A'},
			},
		},
	}

	for _, tc := range testCases {
		got := zigzag.ToMatrix(tc.Test.String, tc.Test.Rows)
		if !cmp.Equal(got, tc.Result) {
			t.Error(cmp.Diff(tc.Result, got))
		}
	}
}

func TestConvert_ConvertsStringIntoItsZigZagString(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Test   Input
		Result string
	}
	testCases := []TestCase{
		{
			Test:   Input{"PAYPALISHIRING", 4},
			Result: "PINALSIGYAHRPI",
		},
		{
			Test:   Input{"PAYPALISHIRING", 3},
			Result: "PAHNAPLSIIGYIR",
		},
		{
			Test:   Input{"A", 1},
			Result: "A",
		},
	}

	for _, tc := range testCases {
		got := zigzag.Convert(tc.Test.String, tc.Test.Rows)
		if !cmp.Equal(got, tc.Result) {
			t.Error(cmp.Diff(tc.Result, got))
		}
	}
}
