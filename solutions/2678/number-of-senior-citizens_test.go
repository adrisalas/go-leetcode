package main

import (
	"testing"
)

type TestCase struct {
	input  []string
	output int
}

func Test(t *testing.T) {
	testCases := []TestCase{
		{[]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}, 2},
		{[]string{"1313579440F2036", "2921522980M5644"}, 0},
	}

	for _, tc := range testCases {
		res := countSeniors(tc.input)
		if res != tc.output {
			t.Errorf("FAIL: input %v  => expected %v, got %v", tc.input, tc.output, res)
		}
	}
}
