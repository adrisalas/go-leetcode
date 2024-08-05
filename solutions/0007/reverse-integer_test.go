package main

import (
	"testing"
)

type TestCase struct {
	input  int
	output int
}

func Test(t *testing.T) {
	testCases := []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for _, tc := range testCases {
		res := reverse(tc.input)
		if res != tc.output {
			t.Errorf("FAIL: input %v  => expected %v, got %v", tc.input, tc.output, res)
		}
	}
}
