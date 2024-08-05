package main

import (
	"testing"
)

type TestCase struct {
	input  string
	output int
}

func Test0004(t *testing.T) {
	testCases := []TestCase{
		{"42", 42},
		{"-042", -42},
		{"0-1", 0},
		{"words and 987", 0},
	}

	for _, tc := range testCases {
		res := myAtoi(tc.input)
		if res != tc.output {
			t.Errorf("FAIL: input %v  => expected %v, got %v", tc.input, tc.output, res)
		}
	}
}
