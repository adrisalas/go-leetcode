package main

import (
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	result []int
}

func Test(t *testing.T) {
	testCases := []TestCase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 1}, 23, []int{-1, -1}},
	}

	for _, tc := range testCases {
		res := twoSum(tc.nums, tc.target)
		unexpectedLength := len(res) != 2
		invalidFirstNumber := res[0] != tc.result[0]
		invalidSecondNumber := res[1] != tc.result[1]
		if unexpectedLength || invalidFirstNumber || invalidSecondNumber {
			t.Errorf("FAIL: input %v, target %d => expected %v, got %v", tc.nums, tc.target, tc.result, res)
		}
	}
}
