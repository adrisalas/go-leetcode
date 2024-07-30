package main

import (
	"testing"
)

type TestCase struct {
	nums1  []int
	nums2  []int
	output float64
}

func Test0004(t *testing.T) {
	testCases := []TestCase{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tc := range testCases {
		res := findMedianSortedArrays(tc.nums1, tc.nums2)
		if res != tc.output {
			t.Errorf("FAIL: nums1 %v, nums2 %v => expected %v, got %v", tc.nums1, tc.nums2, tc.output, res)
		}
	}
}
