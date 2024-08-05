package main

import (
	"testing"
)

type TestCase struct {
	str     string
	numRows int
	output  string
}

func Test(t *testing.T) {
	testCases := []TestCase{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}

	for _, tc := range testCases {
		res := convert(tc.str, tc.numRows)
		if res != tc.output {
			t.Errorf("FAIL: string %v, rows %v  => expected %v, got %v", tc.str, tc.numRows, tc.output, res)
		}
	}
}
