package main

import (
	"testing"
)

type TestCase struct {
	input  string
	output string
}

func Test(t *testing.T) {
	testCases := []TestCase{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, tc := range testCases {
		res := longestPalindrome(tc.input)
		if res != tc.output {
			t.Errorf("FAIL: input %v => expected %v, got %v", tc.input, tc.output, res)
		}
	}
}
