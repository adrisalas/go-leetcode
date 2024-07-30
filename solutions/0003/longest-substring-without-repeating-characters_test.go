package main

import (
	"testing"
)

type TestCase struct {
	input  string
	output int
}

func Test0003(t *testing.T) {
	testCases := []TestCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tc := range testCases {
		res := lengthOfLongestSubstring(tc.input)
		if res != tc.output {
			t.Errorf("FAIL: input %v => expected %v, got %v", tc.input, tc.output, res)
		}
	}
}
